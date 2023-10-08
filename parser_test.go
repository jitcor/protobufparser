package protobufparser

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestProtobufParser_Query(t *testing.T) {
	bin, err := hex.DecodeString("080012001A22108982DA800618BBE8F78B0320F8AFFFD6042889A8AE9B065890ADDCC20A6019700020012A9C1B08DCCBD980061086AFEF1118012288010A290886AFEF1110D8EDAEF90C18A601200B28B9C80130DDCBD9800638F2AD86FB8880808001B801A0CE01120808011000180020001A510A4F0A27080010DDCBD9800618F2AD86FB082000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0332333212044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28BAC80130DACDD9800638CAEDACF88E80808001B801A1CE01120808011000180020001A510A4F0A27080010D9CDD9800618CAEDACF80E2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28BBC80130E5DBD980063890DCC4A38C80808001B801A2CE01120808011000180020001A510A4F0A27080010E5DBD980061890DCC4A30C2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28BCC80130C9DFD9800638D6BDF2F88080808001B801A3CE01120808011000180020001A500A4E0A26080010C8DFD9800618D6BDF2782000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28BDC80130DDDFD9800638D197C0958F80808001B801A4CE01120808011000180020001A510A4F0A27080010DCDFD9800618D197C0950F2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28BEC80130CCE3D9800638D4F180F78180808001B801A5CE01120808011000180020001A510A4F0A27080010CCE3D9800618D4F180F7012000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28BFC80130E0E3D9800638C9EFEAFA8980808001B801A6CE01120808011000180020001A510A4F0A27080010E0E3D9800618C9EFEAFA092000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C0C80130B9E4D9800638958CDDE98C80808001B801A7CE01120808011000180020001A510A4F0A27080010B9E4D9800618958CDDE90C2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C1C80130A6E9D9800638EDE8EBAD8880808001B801A8CE01120808011000180020001A510A4F0A27080010A5E9D9800618EDE8EBAD082000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C2C80130E3EAD9800638ABAEA1A98280808001B801A9CE01120808011000180020001A510A4F0A27080010E4EAD9800618ABAEA1A9022000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28C3C80130F4EAD9800638F0ADCEBC8080808001B801AACE01120808011000180020001A500A4E0A26080010F4EAD9800618F0ADCE3C2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C4C8013092EBD9800638B7D3C5848D80808001B801ABCE01120808011000180020001A510A4F0A2708001092EBD9800618B7D3C5840D2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C5C80130C8EDD9800638B1BDF3ED8980808001B801ACCE01120808011000180020001A510A4F0A27080010C8EDD9800618B1BDF3ED092000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C6C80130B9F3D9800638D1B8C8F08180808001B801ADCE01120808011000180020001A510A4F0A27080010B9F3D9800618D1B8C8F0012000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C7C80130EEF4D9800638DF84CD8E8E80808001B801AECE01120808011000180020001A510A4F0A27080010EEF4D9800618DF84CD8E0E2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C8C8013088F5D9800638F492E6868F80808001B801AFCE01120808011000180020001A510A4F0A2708001088F5D9800618F492E6860F2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28C9C80130ECF6D9800638A980D9FD8C80808001B801B0CE01120808011000180020001A510A4F0A27080010ECF6D9800618A980D9FD0C2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0333333312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28CAC80130A1F8D9800638FAF0D0D48680808001B801B1CE01120808011000180020001A500A4E0A27080010A1F8D9800618FAF0D0D4062000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112060A040A02333312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28CBC80130C5F8D98006388AC7AD9A8A80808001B801B2CE01120808011000180020001A510A4F0A27080010C5F8D98006188AC7AD9A0A2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28CCC80130A7F9D9800638CD9FAD978280808001B801B3CE01120808011000180020001A510A4F0A27080010A7F9D9800618CD9FAD97022000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28CDC801309FFAD9800638E4F1CEC58880808001B801B4CE01120808011000180020001A500A4E0A270800109FFAD9800618E4F1CEC5082000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112060A040A02323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28CEC80130CFFAD9800638B1D4F5EB8780808001B801B5CE01120808011000180020001A500A4E0A27080010CFFAD9800618B1D4F5EB072000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112060A040A02323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28CFC80130E0FBD9800638DACAE5998180808001B801B6CE01120808011000180020001A510A4F0A27080010E0FBD9800618DACAE599012000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002288010A290886AFEF1110D8EDAEF90C18A601200B28D0C80130F4FBD980063898ACB8898E80808001B801B7CE01120808011000180020001A510A4F0A27080010F4FBD980061898ACB8890E2000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112070A050A0331323312044A0208001215AA02129A010FC80100F00100F80100900200CA04002287010A290886AFEF1110D8EDAEF90C18A601200B28D1C8013083FED9800638B08FB5DD8280808001B801B8CE01120808011000180020001A500A4E0A2708001083FED9800618B08FB5DD022000280C300038860140224A0CE5BEAEE8BDAFE99B85E9BB9112060A040A02323312044A0208001215AA02129A010FC80100F00100F80100900200CA04003001380042004800")
	if err != nil {
		t.Error(err)
		return
	}
	value, err := NewProtobufParser(bin).Query(5, 4, 3, 1, 1, 9)
	log.Println("value:", string(value.([]byte)))
}