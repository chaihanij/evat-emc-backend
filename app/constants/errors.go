package constants

var (
	NoLongTextForDecrypt                    = "no long text for decrypt"
	FailToGrabPublicKey                     = "fail to get public key"
	FailUnExpectedSigningMethod             = "fail unexpected signing method"
	FailJwtTokenCannotBeClaimed             = "fail jwt token cannot be claimed"
	NotFoundUserWithGivenAccessToken        = "not found user with given access token"
	GivenUserUuidIsUnmatchedWithJwtUserUuid = "given user uuid is unmatched with jwt user uuid"
	MissingAuthorization                    = "missing authorization header"
	InvalidAuthorization                    = "invalid credential token"
	AuthenticationFailed                    = "authentication failed"
	JWTInvalidStructure                     = "jwt data restore invalid structure"
	JWTRestoreFail                          = "jwt data restore fail"
	UserUIDMissing                          = "user uid is missing"
	DataNotFound                            = "data not found"
	WeakPassword                            = "weak password"
)
