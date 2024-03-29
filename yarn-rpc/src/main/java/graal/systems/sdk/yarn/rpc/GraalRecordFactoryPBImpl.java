package graal.systems.sdk.yarn.rpc;

import com.google.protobuf.ByteString;
import com.google.protobuf.UnknownFieldSet;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.hadoop.classification.InterfaceAudience.Private;
import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.yarn.exceptions.YarnRuntimeException;
import org.apache.hadoop.yarn.factories.RecordFactory;

import java.lang.reflect.Constructor;
import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

@Private
public class GraalRecordFactoryPBImpl implements RecordFactory {
    private static final Log LOG = LogFactory
            .getLog(GraalRecordFactoryPBImpl.class);
    private static final String PB_IMPL_PACKAGE_SUFFIX = "impl.pb";
    private static final String PB_IMPL_CLASS_SUFFIX = "PBImpl";

    private static final GraalRecordFactoryPBImpl self = new GraalRecordFactoryPBImpl();
    private Configuration localConf = new Configuration();
    private ConcurrentMap<Class<?>, Constructor<?>> cache = new ConcurrentHashMap<Class<?>, Constructor<?>>();

    private GraalRecordFactoryPBImpl() {
    }

    public static RecordFactory get() {
        return self;
    }

    @SuppressWarnings("unchecked")
    @Override
    public <T> T newRecordInstance(Class<T> clazz) {

        Constructor<?> constructor = cache.get(clazz);
        String pbImplClassName = getPBImplClassName(clazz);
        if (constructor == null) {
            Class<?> pbClazz;
            try {
                pbClazz = localConf.getClassByName(pbImplClassName);
            } catch (ClassNotFoundException e) {
                throw new YarnRuntimeException("Failed to load class: ["
                        + pbImplClassName + "]", e);
            }
            try {
                constructor = pbClazz.getConstructor();
                constructor.setAccessible(true);
                cache.putIfAbsent(clazz, constructor);
            } catch (NoSuchMethodException e) {
                throw new YarnRuntimeException("Could not find 0 argument constructor", e);
            }
        }

        try {
            Object retObject = constructor.newInstance();

            Class<?> pbClazz = localConf.getClassByName(pbImplClassName);

            LOG.error("Using:class: " + pbClazz);
            if (pbImplClassName.endsWith("RequestPBImpl")) {

                String tenantId = localConf.get("graalsystems.tenant");
                String bridgeId = localConf.get("graalsystems.bridge");
                String projectId = localConf.get("graalsystems.project");
                String infrastructureId = localConf.get("graalsystems.infrastructure");
                String appKey = localConf.get("graalsystems.application.key");
                String appSecret = localConf.get("graalsystems.application.secret");

                String jobId = localConf.get("graalsystems.job");
                String runId = localConf.get("graalsystems.run");

                LOG.error("Using:\ntenant: " + tenantId + "\nbridge: " + bridgeId + "\nproject: " + projectId + "\njob: " + jobId + "\nrun: " + runId + "\ninfrastructure: " + infrastructureId);

                UnknownFieldSet.Builder builder1 = UnknownFieldSet.newBuilder();
                builder1 = builder1.addField(97,
                        UnknownFieldSet.Field.newBuilder()
                                .addLengthDelimited(ByteString.copyFromUtf8(appKey))
                                .build());
                builder1 = builder1.addField(98,
                        UnknownFieldSet.Field.newBuilder()
                                .addLengthDelimited(ByteString.copyFromUtf8(appSecret))
                                .build());

                builder1 = builder1.addField(100,
                        UnknownFieldSet.Field.newBuilder()
                                .addLengthDelimited(ByteString.copyFromUtf8(tenantId))
                                .build());
                builder1 = builder1.addField(101,
                        UnknownFieldSet.Field.newBuilder()
                                .addLengthDelimited(ByteString.copyFromUtf8(bridgeId))
                                .build());

                if (StringUtils.isNotBlank(infrastructureId)) {
                    builder1 = builder1.addField(99,
                            UnknownFieldSet.Field.newBuilder()
                                    .addLengthDelimited(ByteString.copyFromUtf8(infrastructureId))
                                    .build());
                    if (StringUtils.isNotBlank(projectId)) {
                        builder1 = builder1.addField(102,
                                UnknownFieldSet.Field.newBuilder()
                                        .addLengthDelimited(ByteString.copyFromUtf8(projectId))
                                        .build());
                        if (StringUtils.isNotBlank(jobId)) {
                            builder1 = builder1.addField(103,
                                    UnknownFieldSet.Field.newBuilder()
                                            .addLengthDelimited(ByteString.copyFromUtf8(jobId))
                                            .build());
                            if (StringUtils.isNotBlank(runId)) {
                                builder1 = builder1.addField(104,
                                        UnknownFieldSet.Field.newBuilder()
                                                .addLengthDelimited(ByteString.copyFromUtf8(runId))
                                                .build());
                            }
                        }
                    }
                    
                    UnknownFieldSet unknownFields = builder1.build();

                    Field field = pbClazz.getDeclaredField("builder");
                    field.setAccessible(true);
                    Object builder = field.get(retObject);

                    Method setUnknownFields = builder.getClass().getMethod("setUnknownFields", UnknownFieldSet.class);
                    setUnknownFields.setAccessible(true);
                    setUnknownFields.invoke(builder, unknownFields);
                }
            }
            return (T) retObject;
        } catch (InvocationTargetException e) {
            throw new YarnRuntimeException(e);
        } catch (IllegalAccessException e) {
            throw new YarnRuntimeException(e);
        } catch (InstantiationException e) {
            throw new YarnRuntimeException(e);
        } catch (NoSuchFieldException e) {
            throw new YarnRuntimeException(e);
        } catch (NoSuchMethodException e) {
            throw new YarnRuntimeException(e);
        } catch (ClassNotFoundException e) {
            throw new YarnRuntimeException(e);
        }
    }

    private String getPBImplClassName(Class<?> clazz) {
        String srcPackagePart = getPackageName(clazz);
        String srcClassName = getClassName(clazz);
        String destPackagePart = srcPackagePart + "." + PB_IMPL_PACKAGE_SUFFIX;
        String destClassPart = srcClassName + PB_IMPL_CLASS_SUFFIX;
        return destPackagePart + "." + destClassPart;
    }

    private String getClassName(Class<?> clazz) {
        String fqName = clazz.getName();
        return (fqName.substring(fqName.lastIndexOf(".") + 1, fqName.length()));
    }

    private String getPackageName(Class<?> clazz) {
        return clazz.getPackage().getName();
    }
}
