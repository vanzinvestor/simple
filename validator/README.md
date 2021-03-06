# Validator json payload

A package of validators use some validator from [govalidator](https://github.com/asaskevich/govalidator) and add more

## List of validator

- [ ] ByteLength(str string, params ...string)
- [ ] Contains(str, substring string)
- [x] `Equals(str1, str2, errorMessage string)` If Not Equal, then return json error
- [x] `Exists(str, errorMessage string)` If empty, then return json error
- [x] `HasLowerCase(str, errorMessage string)` If has Not lowercase, then return json error
- [x] `HasDigit(str, errorMessage string)` If has Not digit, then return json error
- [x] `HasNotDigit(str, errorMessage string)` If has digit, then return json error
- [x] `HasNotLowerCase(str, errorMessage string)` If has lowercase, then return json error
- [x] `HasNotSpecial(str, errorMessage string)` If has special character, then return json error
- [x] `HasNotUpperCase(str, errorMessage string)` If has uppercase, then return json error
- [x] `HasSpecial(str, errorMessage string)` If has Not special character, then return json error
- [x] `HasNotWhitespace(str, errorMessage string)` If has whitespace, then return json error
- [x] `HasUpperCase(str, errorMessage string)` If has Not uppercase, then return json error
- [x] `HasWhitespace(str, errorMessage string)` If has Not whitespace, then return json error
- [ ] HasWhitespaceOnly(str string)
- [ ] InRange(value interface{}, left interface{}, right interface{})
- [ ] InRangeFloat32(value, left, right float32)
- [ ] InRangeFloat64(value, left, right float64)
- [ ] InRangeInt(value, left, right interface{})
- [ ] IsASCII(str string)
- [ ] IsAlpha(str string)
- [ ] IsAlphanumeric(str string)
- [ ] IsBase64(str string)
- [ ] IsByteLength(str string, min, max int)
- [ ] IsCIDR(str string)
- [ ] IsCRC32(str string)
- [ ] IsCRC32b(str string)
- [ ] IsCreditCard(str string)
- [ ] IsDNSName(str string)
- [ ] IsDataURI(str string)
- [ ] IsDialString(str string)
- [ ] IsDivisibleBy(str, num string)
- [x] `IsEmail(email, errorMessage string)` If Not valid email, then return json error
- [x] `IsEmpty(str, errMessage string)` If Not empty, then return json error
- [ ] IsExistingEmail(email string)
- [ ] IsFilePath(str string) (, int)
- [ ] IsFloat(str string)
- [ ] IsFullWidth(str string)
- [ ] IsHalfWidth(str string)
- [ ] IsHash(str string, algorithm string)
- [ ] IsHexadecimal(str string)
- [ ] IsHexcolor(str string)
- [ ] IsHost(str string)
- [ ] IsIP(str string)
- [ ] IsIPv4(str string)
- [ ] IsIPv6(str string)
- [ ] IsISBN(str string, version int)
- [ ] IsISBN10(str string)
- [ ] IsISBN13(str string)
- [ ] IsISO3166Alpha2(str string)
- [ ] IsISO3166Alpha3(str string)
- [ ] IsISO4217(str string)
- [ ] IsISO693Alpha2(str string)
- [ ] IsISO693Alpha3b(str string)
- [ ] IsIn(str string, params ...string)
- [ ] IsInRaw(str string, params ...string)
- [ ] IsInt(str string)
- [ ] IsJSON(str string)
- [ ] IsLatitude(str string)
- [ ] IsLongitude(str string)
- [x] `IsLowerCase(str, errorMessage string)` If Not lowercase, then return json error
- [ ] IsMAC(str string)
- [ ] IsMD4(str string)
- [ ] IsMD5(str string)
- [ ] IsMagnetURI(str string)
- [x] `IsLengthMax(str string, max int, errorMessage string)` If string length more than maximum, then return json error
- [x] `IsLengthMin(str string, min int, errorMessage string)` If string length less than minimum, then return json error
- [ ] IsMongoID(str string)
- [ ] IsMultibyte(str string)
- [ ] IsNatural(value float64)
- [ ] IsNegative(value float64)
- [ ] IsNonNegative(value float64)
- [ ] IsNonPositive(value float64)
- [x] `IsNotEmpty(str, errMessage string)` If empty, then return json error
- [ ] IsNotNull(str string)
- [ ] IsNull(str string)
- [ ] IsNumeric(str string)
- [ ] IsPort(str string)
- [ ] IsPositive(value float64)
- [ ] IsPrintableASCII(str string)
- [ ] IsRFC3339(str string)
- [ ] IsRFC3339WithoutZone(str string)
- [ ] IsRGBcolor(str string)
- [ ] IsRequestURI(rawurl string)
- [ ] IsRequestURL(rawurl string)
- [ ] IsRipeMD128(str string)
- [ ] IsRipeMD160(str string)
- [ ] IsRsaPub(str string, params ...string)
- [ ] IsRsaPublicKey(str string, keylen int)
- [ ] IsSHA1(str string)
- [ ] IsSHA256(str string)
- [ ] IsSHA384(str string)
- [ ] IsSHA512(str string)
- [ ] IsSSN(str string)
- [ ] IsSemver(str string)
- [ ] IsTiger128(str string)
- [ ] IsTiger160(str string)
- [ ] IsTiger192(str string)
- [ ] IsTime(str string, format string)
- [ ] IsType(v interface{}, params ...string)
- [ ] `IsURL(str, errorMessage string)` If Not URL format , then return json error
- [ ] IsUTFDigit(str string)
- [ ] IsUTFLetter(str string)
- [ ] IsUTFLetterNumeric(str string)
- [ ] IsUTFNumeric(str string)
- [ ] IsUUID(str string)
- [ ] IsUUIDv3(str string)
- [ ] IsUUIDv4(str string)
- [ ] IsUUIDv5(str string)
- [ ] IsUnixTime(str string)
- [x] `IsUpperCase(str, errorMessage string)` If Not uppercase , then return json error
- [ ] IsVariableWidth(str string)
- [ ] IsWhole(value float64)
- [x] `Matches(regex, str, errorMessage string)` If Not Matched, then return json error
- [ ] MaxStringLength(str string, params ...string)
- [ ] MinStringLength(str string, params ...string)
- [x] `NotMatches(regex, str, errorMessage string)` If Matched, then return json error
- [x] `NotEquals(str1, str2, errorMessage string)` If Equal, then return json error
- [ ] Range(str string, params ...string)
- [ ] Range(str string, params ...string)
- [ ] RuneLength(str string, params ...string)
- [ ] SetFieldsRequiredByDefault(value )
- [ ] SetNilPtrAllowedByRequired(value )
- [ ] StringLength(str string, params ...string)
- [ ] StringMatches(s string, params ...string)
