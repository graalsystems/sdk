package graal.systems.sdk.parquet.kms;


import org.apache.hadoop.conf.Configuration;
import org.apache.http.HttpEntity;
import org.apache.http.client.config.RequestConfig;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;
import org.apache.parquet.crypto.KeyAccessDeniedException;
import org.apache.parquet.crypto.ParquetCryptoRuntimeException;
import org.apache.parquet.crypto.keytools.KeyToolkit;
import org.apache.parquet.crypto.keytools.KmsClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.charset.StandardCharsets;

public class InternalKmsClient implements KmsClient {

    private static final Logger LOGGER = LoggerFactory.getLogger(InternalKmsClient.class);

    private CloseableHttpClient client;
    private String kmsInstanceID;
    private String kmsInstanceURL;
    private String accessToken;

    @Override
    public synchronized void initialize(Configuration configuration, String kmsInstanceID, String kmsInstanceURL, String accessToken) {

        RequestConfig requestConfig = RequestConfig.custom()
                .setConnectionRequestTimeout(5000)
                .setConnectTimeout(5000)
                .setSocketTimeout(5000)
                .build();

        this.client = HttpClientBuilder.create()
                .setDefaultRequestConfig(requestConfig)
                .build();

        this.kmsInstanceID = kmsInstanceID;
        this.kmsInstanceURL = kmsInstanceURL;
        this.accessToken = accessToken;
    }

    @Override
    public synchronized String wrapKey(byte[] keyBytes, String masterKeyIdentifier)
            throws KeyAccessDeniedException, UnsupportedOperationException {
        byte[] masterKey = getMasterKey(masterKeyIdentifier);
        byte[] AAD = masterKeyIdentifier.getBytes(StandardCharsets.UTF_8);
        return KeyToolkit.encryptKeyLocally(keyBytes, masterKey, AAD);
    }

    private byte[] getMasterKey(String masterKeyIdentifier) {
        byte[] masterKey;
        HttpGet get = new HttpGet(this.kmsInstanceURL + masterKeyIdentifier + "?data");
        get.addHeader("X-Access-Token", this.accessToken);
        get.addHeader("X-Tenant", this.kmsInstanceID);
        try (CloseableHttpResponse response = this.client.execute(get)) {
            HttpEntity entity = response.getEntity();
            if (entity != null) {
                masterKey = EntityUtils.toByteArray(entity);
            } else {
                throw new ParquetCryptoRuntimeException("Key not found: " + masterKeyIdentifier);
            }
        } catch (Exception e) {
            throw new ParquetCryptoRuntimeException("Error while retrieving key: " + masterKeyIdentifier, e);
        }
        return masterKey;
    }

    @Override
    public synchronized byte[] unwrapKey(String wrappedKey, String masterKeyIdentifier)
            throws KeyAccessDeniedException, UnsupportedOperationException {
        byte[] masterKey = getMasterKey(masterKeyIdentifier);
        byte[] AAD = masterKeyIdentifier.getBytes(StandardCharsets.UTF_8);
        return KeyToolkit.decryptKeyLocally(wrappedKey, masterKey, AAD);
    }
}