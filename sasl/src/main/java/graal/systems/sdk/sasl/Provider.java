package graal.systems.sdk.sasl;

import graal.systems.sdk.sasl.digest.FactoryImpl;

import java.security.*;

/**
 * The SASL provider.
 * Provides client support for
 * - DIGEST-MD5
 * And server support for
 * - DIGEST-MD5
 */

public final class Provider extends java.security.Provider {

    private static final long serialVersionUID = 8622598936488630849L;

    private static final String info = "Custom SASL provider" +
            "(implements client mechanisms for: " +
            "DIGEST-MD5;" +
            " server mechanisms for: DIGEST-MD5)";

    private static final class ProviderService
            extends java.security.Provider.Service {
        ProviderService(java.security.Provider p, String type, String algo,
                        String cn) {
            super(p, type, algo, cn, null, null);
        }

        @Override
        public Object newInstance(Object ctrParamObj)
                throws NoSuchAlgorithmException {
            String type = getType();
            if (ctrParamObj != null) {
                throw new InvalidParameterException
                        ("constructorParameter not used with " + type + " engines");
            }

            String algo = getAlgorithm();
            try {
                // DIGEST-MD5, NTLM uses same impl class for client and server
                if (algo.equals("DIGEST-MD5")) {
                    return new FactoryImpl();
                }
            } catch (Exception ex) {
                throw new NoSuchAlgorithmException("Error constructing " +
                        type + " for " + algo + " using InternalSasl", ex);
            }
            throw new ProviderException("No impl for " + algo +
                    " " + type);
        }
    }

    public Provider() {
        super("InternalSASL", "0.0.1", info);

        final Provider p = this;
        AccessController.doPrivileged(new PrivilegedAction<Void>() {
            public Void run() {
                // Client mechanisms
                putService(new Provider.ProviderService(p, "SaslClientFactory",
                        "DIGEST-MD5", "graal.systems.sdk.sasl.FactoryImpl"));

                // Server mechanisms
                putService(new Provider.ProviderService(p, "SaslServerFactory",
                        "DIGEST-MD5", "graal.systems.sdk.sasl.FactoryImpl"));
                return null;
            }
        });
    }
}
