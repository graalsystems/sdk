package graal.systems.sdk.yarn.rpc;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.yarn.api.protocolrecords.RegisterApplicationMasterRequest;
import org.apache.hadoop.yarn.api.protocolrecords.impl.pb.RegisterApplicationMasterRequestPBImpl;
import org.apache.hadoop.yarn.factories.RecordFactory;
import org.apache.hadoop.yarn.factory.providers.RecordFactoryProvider;
import org.junit.jupiter.api.Test;

public class GraalRecordFactoryPBImplTests {

    @Test
    public void test() {
        Configuration conf = new Configuration(false);
        conf.set("yarn.ipc.record.factory.class", "graal.systems.sdk.yarn.rpc.GraalRecordFactoryPBImpl");
        RecordFactory recordFactory = RecordFactoryProvider.getRecordFactory(conf);

        RegisterApplicationMasterRequest registerApplicationMasterRequest = recordFactory.newRecordInstance(RegisterApplicationMasterRequest.class);

        System.err.println(((RegisterApplicationMasterRequestPBImpl) registerApplicationMasterRequest).getProto().getUnknownFields());
    }
}
