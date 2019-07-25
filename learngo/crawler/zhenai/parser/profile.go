package parser

const  ageRe = `<div class="m-btn purple" data-v-bff6f798="">29岁</div>`
//func ParseProfile(contents []byte) engine.ParserResult {
//	profile := model.Profile{}
//	re := regexp.MustCompile(ageRe)
//	match := re.FindSubmatch(contents)
//	if match != nil {
//		age, err := strconv.Atoi(string[match[1]])
//		if err != nil {
//			profile.Age = age
//		}
//	}
//	return
//}