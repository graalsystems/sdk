package graal.systems.sdk.yarn.rpc;

import com.google.protobuf.ByteString;
import com.google.protobuf.UnknownFieldSet;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.yarn.api.protocolrecords.RegisterApplicationMasterRequest;
import org.apache.hadoop.yarn.api.protocolrecords.RegisterApplicationMasterResponse;
import org.apache.hadoop.yarn.api.protocolrecords.impl.pb.RegisterApplicationMasterRequestPBImpl;
import org.apache.hadoop.yarn.api.protocolrecords.impl.pb.RegisterApplicationMasterResponsePBImpl;
import org.apache.hadoop.yarn.factories.RecordFactory;
import org.apache.hadoop.yarn.factory.providers.RecordFactoryProvider;
import org.junit.jupiter.api.Test;

import static org.assertj.core.api.Assertions.assertThat;
import static uk.org.webcompere.systemstubs.SystemStubs.restoreSystemProperties;
import static uk.org.webcompere.systemstubs.SystemStubs.withEnvironmentVariable;

public class GraalRecordFactoryPBImplTests {

    @Test
    public void withEnvForARequest() throws Exception {
        restoreSystemProperties(() -> {
            withEnvironmentVariable("GRAAL_TENANT", "tenant-1")
                    .execute(() -> {

                        // Given
                        Configuration conf = new Configuration(false);
                        conf.set("yarn.ipc.record.factory.class", "graal.systems.sdk.yarn.rpc.GraalRecordFactoryPBImpl");
                        RecordFactory recordFactory = RecordFactoryProvider.getRecordFactory(conf);

                        // When
                        RegisterApplicationMasterRequest registerApplicationMasterRequest = recordFactory.newRecordInstance(RegisterApplicationMasterRequest.class);

                        // Then
                        UnknownFieldSet unknownFields = ((RegisterApplicationMasterRequestPBImpl) registerApplicationMasterRequest).getProto().getUnknownFields();

                        assertThat(unknownFields.asMap())
                                .hasSize(1)
                                .containsEntry(100,
                                        UnknownFieldSet.Field.newBuilder()
                                                .addLengthDelimited(ByteString.copyFromUtf8("tenant-1"))
                                                .build());
                    });
        });
    }

    @Test
    public void withEnvForNotARequest() throws Exception {
        restoreSystemProperties(() -> {
            withEnvironmentVariable("GRAAL_TENANT", "tenant-2")
                    .execute(() -> {

                        // Given
                        Configuration conf = new Configuration(false);
                        conf.set("yarn.ipc.record.factory.class", "graal.systems.sdk.yarn.rpc.GraalRecordFactoryPBImpl");
                        RecordFactory recordFactory = RecordFactoryProvider.getRecordFactory(conf);

                        // When
                        RegisterApplicationMasterResponse registerApplicationMasterResponse = recordFactory.newRecordInstance(RegisterApplicationMasterResponse.class);

                        // Then
                        UnknownFieldSet unknownFields = ((RegisterApplicationMasterResponsePBImpl) registerApplicationMasterResponse).getProto().getUnknownFields();

                        assertThat(unknownFields.asMap()).hasSize(0);
                    });
        });
    }
}
