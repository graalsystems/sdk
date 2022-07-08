package graal.systems.sdk.sasl.digest;

import graal.systems.sdk.sasl.util.PolicyUtils;

import javax.security.auth.callback.CallbackHandler;
import javax.security.sasl.*;
import java.util.Map;


/**
 * Client and server factory for DIGEST-MD5 SASL client/server mechanisms.
 * See DigestMD5Client and DigestMD5Server for input requirements.
 *
 * @author Jonathan Bruce
 * @author Rosanna Lee
 */

public final class FactoryImpl implements SaslClientFactory,
        SaslServerFactory {

    private static final String[] myMechs = {"DIGEST-MD5"};
    private static final int DIGEST_MD5 = 0;
    private static final int[] mechPolicies = {
            PolicyUtils.NOPLAINTEXT | PolicyUtils.NOANONYMOUS};

    /**
     * Empty constructor.
     */
    public FactoryImpl() {
    }

    /**
     * Returns a new instance of the DIGEST-MD5 SASL client mechanism.
     *
     * @return a new SaslClient; otherwise null if unsuccessful.
     * @throws SaslException If there is an error creating the DigestMD5
     *                       SASL client.
     */
    public SaslClient createSaslClient(String[] mechs,
                                       String authorizationId, String protocol, String serverName,
                                       Map<String, ?> props, CallbackHandler cbh)
            throws SaslException {

        for (int i = 0; i < mechs.length; i++) {
            if (mechs[i].equals(myMechs[DIGEST_MD5]) &&
                    PolicyUtils.checkPolicy(mechPolicies[DIGEST_MD5], props)) {

                if (cbh == null) {
                    throw new SaslException(
                            "Callback handler with support for RealmChoiceCallback, " +
                                    "RealmCallback, NameCallback, and PasswordCallback " +
                                    "required");
                }

                return new DigestMD5Client(authorizationId,
                        protocol, serverName, props, cbh);
            }
        }
        return null;
    }

    /**
     * Returns a new instance of the DIGEST-MD5 SASL server mechanism.
     *
     * @return a new SaslServer; otherwise null if unsuccessful.
     * @throws SaslException If there is an error creating the DigestMD5
     *                       SASL server.
     */
    public SaslServer createSaslServer(String mech,
                                       String protocol, String serverName, Map<String, ?> props, CallbackHandler cbh)
            throws SaslException {

        if (mech.equals(myMechs[DIGEST_MD5]) &&
                PolicyUtils.checkPolicy(mechPolicies[DIGEST_MD5], props)) {

            if (cbh == null) {
                throw new SaslException(
                        "Callback handler with support for AuthorizeCallback, " +
                                "RealmCallback, NameCallback, and PasswordCallback " +
                                "required");
            }

            return new DigestMD5Server(protocol, serverName, props, cbh);
        }
        return null;
    }

    /**
     * Returns the authentication mechanisms that this factory can produce.
     *
     * @return String[] {"DigestMD5"} if policies in env match those of this
     * factory.
     */
    public String[] getMechanismNames(Map<String, ?> env) {
        return PolicyUtils.filterMechs(myMechs, mechPolicies, env);
    }
}
