package oidc

import (
	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

// NewOpenIDConnectWellKnownConfiguration generates a new OpenIDConnectWellKnownConfiguration.
func NewOpenIDConnectWellKnownConfiguration(c *schema.OpenIDConnectConfiguration) (config OpenIDConnectWellKnownConfiguration) {
	config = OpenIDConnectWellKnownConfiguration{
		OAuth2WellKnownConfiguration: OAuth2WellKnownConfiguration{
			CommonDiscoveryOptions: CommonDiscoveryOptions{
				SubjectTypesSupported: []string{
					SubjectTypePublic,
					SubjectTypePairwise,
				},
				ResponseTypesSupported: []string{
					ResponseTypeAuthorizationCodeFlow,
					ResponseTypeImplicitFlowIDToken,
					ResponseTypeImplicitFlowToken,
					ResponseTypeImplicitFlowBoth,
					ResponseTypeHybridFlowIDToken,
					ResponseTypeHybridFlowToken,
					ResponseTypeHybridFlowBoth,
				},
				GrantTypesSupported: []string{
					GrantTypeAuthorizationCode,
					GrantTypeImplicit,
					GrantTypeRefreshToken,
				},
				ResponseModesSupported: []string{
					ResponseModeFormPost,
					ResponseModeQuery,
					ResponseModeFragment,
				},
				ScopesSupported: []string{
					ScopeOfflineAccess,
					ScopeOpenID,
					ScopeProfile,
					ScopeGroups,
					ScopeEmail,
				},
				ClaimsSupported: []string{
					ClaimAuthenticationMethodsReference,
					ClaimAudience,
					ClaimAuthorizedParty,
					ClaimClientIdentifier,
					ClaimExpirationTime,
					ClaimIssuedAt,
					ClaimIssuer,
					ClaimJWTID,
					ClaimRequestedAt,
					ClaimSubject,
					ClaimAuthenticationTime,
					ClaimNonce,
					ClaimPreferredEmail,
					ClaimEmailVerified,
					ClaimEmailAlts,
					ClaimGroups,
					ClaimPreferredUsername,
					ClaimFullName,
				},
				TokenEndpointAuthMethodsSupported: []string{
					ClientAuthMethodClientSecretBasic,
					ClientAuthMethodClientSecretPost,
					ClientAuthMethodNone,
				},
			},
			OAuth2DiscoveryOptions: OAuth2DiscoveryOptions{
				CodeChallengeMethodsSupported: []string{
					PKCEChallengeMethodSHA256,
				},
			},
			OAuth2PushedAuthorizationDiscoveryOptions: &OAuth2PushedAuthorizationDiscoveryOptions{
				RequirePushedAuthorizationRequests: c.PAR.Enforce,
			},
		},

		OpenIDConnectDiscoveryOptions: OpenIDConnectDiscoveryOptions{
			IDTokenSigningAlgValuesSupported: []string{
				SigningAlgorithmRSAWithSHA256,
			},
			UserinfoSigningAlgValuesSupported: []string{
				SigningAlgorithmNone,
				SigningAlgorithmRSAWithSHA256,
			},
		},
		OpenIDConnectFrontChannelLogoutDiscoveryOptions: &OpenIDConnectFrontChannelLogoutDiscoveryOptions{},
		OpenIDConnectBackChannelLogoutDiscoveryOptions:  &OpenIDConnectBackChannelLogoutDiscoveryOptions{},
		OpenIDConnectPromptCreateDiscoveryOptions: &OpenIDConnectPromptCreateDiscoveryOptions{
			PromptValuesSupported: []string{
				PromptNone,
				PromptConsent,
			},
		},
	}

	if c.EnablePKCEPlainChallenge {
		config.CodeChallengeMethodsSupported = append(config.CodeChallengeMethodsSupported, PKCEChallengeMethodPlain)
	}

	return config
}

// Copy the values of the OAuth2WellKnownConfiguration and return it as a new struct.
func (opts OAuth2WellKnownConfiguration) Copy() (optsCopy OAuth2WellKnownConfiguration) {
	optsCopy = OAuth2WellKnownConfiguration{
		CommonDiscoveryOptions: opts.CommonDiscoveryOptions,
		OAuth2DiscoveryOptions: opts.OAuth2DiscoveryOptions,
	}

	if opts.OAuth2DeviceAuthorizationGrantDiscoveryOptions != nil {
		optsCopy.OAuth2DeviceAuthorizationGrantDiscoveryOptions = &OAuth2DeviceAuthorizationGrantDiscoveryOptions{}
		*optsCopy.OAuth2DeviceAuthorizationGrantDiscoveryOptions = *opts.OAuth2DeviceAuthorizationGrantDiscoveryOptions
	}

	if opts.OAuth2MutualTLSClientAuthenticationDiscoveryOptions != nil {
		optsCopy.OAuth2MutualTLSClientAuthenticationDiscoveryOptions = &OAuth2MutualTLSClientAuthenticationDiscoveryOptions{}
		*optsCopy.OAuth2MutualTLSClientAuthenticationDiscoveryOptions = *opts.OAuth2MutualTLSClientAuthenticationDiscoveryOptions
	}

	if opts.OAuth2IssuerIdentificationDiscoveryOptions != nil {
		optsCopy.OAuth2IssuerIdentificationDiscoveryOptions = &OAuth2IssuerIdentificationDiscoveryOptions{}
		*optsCopy.OAuth2IssuerIdentificationDiscoveryOptions = *opts.OAuth2IssuerIdentificationDiscoveryOptions
	}

	if opts.OAuth2JWTIntrospectionResponseDiscoveryOptions != nil {
		optsCopy.OAuth2JWTIntrospectionResponseDiscoveryOptions = &OAuth2JWTIntrospectionResponseDiscoveryOptions{}
		*optsCopy.OAuth2JWTIntrospectionResponseDiscoveryOptions = *opts.OAuth2JWTIntrospectionResponseDiscoveryOptions
	}

	if opts.OAuth2JWTSecuredAuthorizationRequestDiscoveryOptions != nil {
		optsCopy.OAuth2JWTSecuredAuthorizationRequestDiscoveryOptions = &OAuth2JWTSecuredAuthorizationRequestDiscoveryOptions{}
		*optsCopy.OAuth2JWTSecuredAuthorizationRequestDiscoveryOptions = *opts.OAuth2JWTSecuredAuthorizationRequestDiscoveryOptions
	}

	if opts.OAuth2PushedAuthorizationDiscoveryOptions != nil {
		optsCopy.OAuth2PushedAuthorizationDiscoveryOptions = &OAuth2PushedAuthorizationDiscoveryOptions{}
		*optsCopy.OAuth2PushedAuthorizationDiscoveryOptions = *opts.OAuth2PushedAuthorizationDiscoveryOptions
	}

	return optsCopy
}

// Copy the values of the OpenIDConnectWellKnownConfiguration and return it as a new struct.
func (opts OpenIDConnectWellKnownConfiguration) Copy() (optsCopy OpenIDConnectWellKnownConfiguration) {
	optsCopy = OpenIDConnectWellKnownConfiguration{
		OAuth2WellKnownConfiguration:  opts.OAuth2WellKnownConfiguration.Copy(),
		OpenIDConnectDiscoveryOptions: opts.OpenIDConnectDiscoveryOptions,
	}

	if opts.OpenIDConnectFrontChannelLogoutDiscoveryOptions != nil {
		optsCopy.OpenIDConnectFrontChannelLogoutDiscoveryOptions = &OpenIDConnectFrontChannelLogoutDiscoveryOptions{}
		*optsCopy.OpenIDConnectFrontChannelLogoutDiscoveryOptions = *opts.OpenIDConnectFrontChannelLogoutDiscoveryOptions
	}

	if opts.OpenIDConnectBackChannelLogoutDiscoveryOptions != nil {
		optsCopy.OpenIDConnectBackChannelLogoutDiscoveryOptions = &OpenIDConnectBackChannelLogoutDiscoveryOptions{}
		*optsCopy.OpenIDConnectBackChannelLogoutDiscoveryOptions = *opts.OpenIDConnectBackChannelLogoutDiscoveryOptions
	}

	if opts.OpenIDConnectSessionManagementDiscoveryOptions != nil {
		optsCopy.OpenIDConnectSessionManagementDiscoveryOptions = &OpenIDConnectSessionManagementDiscoveryOptions{}
		*optsCopy.OpenIDConnectSessionManagementDiscoveryOptions = *opts.OpenIDConnectSessionManagementDiscoveryOptions
	}

	if opts.OpenIDConnectRPInitiatedLogoutDiscoveryOptions != nil {
		optsCopy.OpenIDConnectRPInitiatedLogoutDiscoveryOptions = &OpenIDConnectRPInitiatedLogoutDiscoveryOptions{}
		*optsCopy.OpenIDConnectRPInitiatedLogoutDiscoveryOptions = *opts.OpenIDConnectRPInitiatedLogoutDiscoveryOptions
	}

	if opts.OpenIDConnectPromptCreateDiscoveryOptions != nil {
		optsCopy.OpenIDConnectPromptCreateDiscoveryOptions = &OpenIDConnectPromptCreateDiscoveryOptions{}
		*optsCopy.OpenIDConnectPromptCreateDiscoveryOptions = *opts.OpenIDConnectPromptCreateDiscoveryOptions
	}

	if opts.OpenIDConnectClientInitiatedBackChannelAuthFlowDiscoveryOptions != nil {
		optsCopy.OpenIDConnectClientInitiatedBackChannelAuthFlowDiscoveryOptions = &OpenIDConnectClientInitiatedBackChannelAuthFlowDiscoveryOptions{}
		*optsCopy.OpenIDConnectClientInitiatedBackChannelAuthFlowDiscoveryOptions = *opts.OpenIDConnectClientInitiatedBackChannelAuthFlowDiscoveryOptions
	}

	if opts.OpenIDConnectJWTSecuredAuthorizationResponseModeDiscoveryOptions != nil {
		optsCopy.OpenIDConnectJWTSecuredAuthorizationResponseModeDiscoveryOptions = &OpenIDConnectJWTSecuredAuthorizationResponseModeDiscoveryOptions{}
		*optsCopy.OpenIDConnectJWTSecuredAuthorizationResponseModeDiscoveryOptions = *opts.OpenIDConnectJWTSecuredAuthorizationResponseModeDiscoveryOptions
	}

	if opts.OpenIDFederationDiscoveryOptions != nil {
		optsCopy.OpenIDFederationDiscoveryOptions = &OpenIDFederationDiscoveryOptions{}
		*optsCopy.OpenIDFederationDiscoveryOptions = *opts.OpenIDFederationDiscoveryOptions
	}

	return optsCopy
}
