flash:
	tinygo flash -size short -target badger2040 -ldflags="-X main.YourName='$(NAME)' \
	-X main.YourTitleA1='$(TITLEA1)' -X main.YourTitleA2='$(TITLEA2)'  -X main.YourTitleB1='$(TITLEB1)' \
	 -X main.YourTitleB2='$(TITLEB2)' -X main.YourMarqueeTop='$(MARQUEETOP)' -X main.YourMarqueeMiddle='$(MARQUEEMIDDLE)' \
	 -X main.YourMarqueeBottom='$(MARQUEEBOTTOM)' -X main.YourQRText='$(QRTEXT)'" .

ron:
	tinygo flash --target=badger2040 -ldflags="-X main.YourName='RON' \
-X main.YourCompany='Hybrid Group' \
-X main.YourTitle='Technologist for hire' \
-X main.YourSocial='@deadprogram' \
-X main.ProfilePic='AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA4OAAAAAAAAAAAAAAAAAA5//+AAAAAAAAAAAAAAAAMPf//ngAAAAAAAAAAAAACD/////5wAAAAAAAAAAAA59//////+EAAAAAAAAAAAOf///////vAAAAAAAAAAAz/////////wwAAAAAAAAA///////////8AAAAAAAAAP///////////+AAAAAAABf////////////wAAAAAAB/////////////8AAAAAAAf/////////////+AAAAAAP//////////////4AAAAAf/////////+////+AAAAAB///////+///////gAAAAA///////+/+//////AAAAD///////9/////3//AAAAAf//3//////////wD4AAAAH/Af//////7///4AP4AAAD/gD//////////8AD8AAAH/wAf////D/8//+CA+AAAH/8CD/++fACAH9/AgP4AAB//AX///AAQAED//YD/wAAB/wf///gAgkQH//+Y/+AAAf+f3/v4giABP///+f/gAAP//7+GPCAJIHwf/v3+AAAf//j/+A8EgATwBn4//gAAP//A//+HgCSB4///h/8AAAH/4/8H45kAI4f4f9//4AAH///8APjASQcfAB///+AAP///+AA84gBHPAAP//4AAB////gADniSHngAD///gAAB///wAAY4ALxgAA///+AAA///8AADPEg44AAD///AAAf//8AAA54CecAAB//+AAAH///gAAGeIGOAAAf//wAAAf//4AAAxgljAAAH//+AAAP///AAAMZA4wAAB//+AAAf//9gAABnEsYAAAH//gAAH//nwAAAYwDGAAAAe/8AAAD//AAAAGMkxgAAAH//wAAA//gAAAAjAcYAAAA/84AAAP/4AAAAO0mcAYAAP/sAAAD/4AAAeDv+xAHgAB//AAAAn+AAALwj/5wI4AAD+wAAAP+AAAEcO4fcH+AAA/sAAAD4AAABPDshhh/gAAH7AAAAnAAAA/wjKZwP4AAAOwAAANwAAAH8IwLED8AAADsAAADMAAAB/GcQxgMAAAA7AAAAzAAAAPhndN4AAAAAOwAAAMwAAAAAR//GAAAAADIAAABOAAAAAEfP4wAAAAB2AAAAbgAAAADffeMAAAAAdgAAAGcAAAAAn9f7gAAAAGYAAABnAAAAAb/3eYAAAADsAAAAMwAAAAM8/jzAAAAAzAAAADOAAAADeTA84AAAAdgAAAAbwAAADHsBHmAAAAOYAAAAHuAAABziCEc4AAADsAAAAA74AAAx5IIHHAAAB/AAAAAGfAAA4+W8n88AAB7gAAAABx4AA487f/njwAB8wAAAAAOH4L4eP5Hg8fwB+cAAAAAHwf/gfAWRADw//8OAAAAABPgSAfBBEYQ+AvoHgAAAAAR+AD/hGYEhB8AAfsAAAAAGP+//AAGBAEP/J/yAAAAABIP/+AiE2xgIf//AwAAAAAQAdICCIH5CIAP+AMAAAAAGRAAEIAIAAISAACJAAAAABAEJAQkYgJgQIACAwAAAAASYQEBAQCQCQgSSGMAAAAAGAhIYEgQDQAiAAEJAAAAABECAggCBH/kgIkkAwAAAAAQYJCCIIJAcBAAAIMAAAAAGAgEEIgQwBJEZJITAAAAABICQQQCBDAYAQAAAwAAAAAQgBBhIIEMCSASSQkAAAAAEBEECAgQBGgIgABjAAAAABkEQQJCAAQIggkkAwAAAAAQYBBgEAAJCCBgAREAAAAAEgkECQAPmBkJBJAFAAAAABiAQQBABfAMQBAEgwAAAAAQEhBkAARhBhICQBMAAAAAEgCEAAEEAGMAiBIBAAAAABCIIRIABAQAiAEAiQAAAAAYAggABAIAiMJnyAMAAAAAEmCiYAADLmBh/uJjAAAAABAIf/QQEPESH5AgCQAAAAARAkSfgAAAiLwAMQEAAAAAGGCAAeAAAGTgABhlAAAAABIJgAAwAAATgHsYAwAAAAAQgZsAGAAACgAACREAAAAAEBEACAhgAAYCBAgFAAAAABkFBAIIAQAGCABMgQAAAAAQYQBgaAAEAkACGBMAAAAAEgkhBJgAAGMkiIkBAAAAABiBkHJwBAABknf4ZQAAAAAQEf+P4GAgAPv/2AMAAAAAEgTv/gAAAAA/gIIRAAAAABiAAAAAAACAAAiAhQAAAAAQEgAAAAACAARoiAMAAAAAEgCQAAABAAIDASJjAAAAABiIAAABBAAAAQYACQAAAAAQAkgABAAIAADkkQMAAAAAEmAAACAAAAgQGARjAAAAABgJIQkAEEEAgAEgCQAAAAAQAAAAAAAAAAAAAAEAAA='" --monitor --stack-size=8kb .

conejo:
	tinygo flash --target=badger2040 -ldflags="-X main.YourName='CONEJO' \
-X main.YourCompany='TinyGo' \
-X main.YourTitle='Technologist for hire' \
-X main.YourSocial='@conejo@social.tinygo.org' \
-X main.ProfilePic='/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////83f///////////////////2Zl//////////////////tjuaPt////////////////nb79Zsf//////////////P3Gx5r7P/////////////9me3r7Odv/////////////p7tezd7N////////////9uWt3bbe9b////////////bd22W24zan///////////vnna+2z3b4///////////+2Otutve3Xz//////////5t9u2Ns222ff/////////+23dbdt3u243/////////+7bt23bdm3v2//////////22mbbts3dsd/////////+NbZ7dm259t5m/////////623m13dtjbuef////////nu2e3Xtt/bM59p///////t2259tntp293mbn//////5nNuVu3du7bd+7Zf/////5uds/bf5tZu55u13/////+777+/cnt2bb3s3e//////7mrZ2e+drfdud247/////23bb9+99u2bc92z1/////9bt3u22+7b7b+23d3////+273fu3ZrWm3t7t3v/////+/eAG+3u3ffbdtmM/////154AAN7e7t9t53e93/////X4AAAz53vjvb7Zve/////uYAAADL3bPue732c/////28AAAAY7t/t+52Pd9/////8AAAADBu2722WA+9/////lAAAAAQHe3tveAGbf////3gAAAAGYdvt8+ABnf/////oAAAAAwhu954AA/d/////sAAAAAMADp7gABH7f////9AAAAABEgAIAAAT5v/////gAAAAAYCRGAAAP+N/////4AAAAAGEBBgAAB/gf////+AAAAAAgSCYAAAf4H/////gfAAAAIgIEAAAD4A/////wJ4AAADiQnAAAAAAP////8EPAAAAgBAYAAAAAD/////BjwAAAIkEkAAAAAA/////w/4AAADgQHAAAAAAP////+H+AAAAhBIYAAAAAD/////A/gAAAIEAkAAAAAI/////4HwAAAGISBgAAAAGf/////AAAAABg9J4AAAAHX/////8AAAAARf8CAAAAHd/////5gAAAAEPfIwAAAGT//////sAAAADHt4cAAABnb/////ZgAAAAh7fBgAAAXb/////5vAAAAbvu+YAAAN2/////77oAAAPc/r9AAAC7f/////zfgAAc3bNjeAAA5t/////nZMAA5zttXN4AA13v/////btwAztm3d+zAA3db////927bB3b3bbjLcA2Z7/////m3a/mbN27Pe2417rv/////t2yZu9n/922/2ze3////9tm3123vJH7dsW7Z3/////vZtnd2bmRPN99t2X//////btvewADEYAB3t29n/////27bbgiQYECIGNpu//////3ZtukCAWBGIm9vt5/////+2ba5kEg2wADvbbH//////vdtt0QAH5ETmNneZ/////+3bd7xJkAATne3b///////vttXnIBbCLXe7bGf/////9m29m75Nm+1tzu/9/////93tq/rT7X2brHW2N//////3tu5O3bPn8ve3d9//////27bZ9bY/Ad7Xu1vr/////93bd7drzQBnOc3ce//////N22ZbbfPCee5279//////x+272t2+cD7Wdru3/////8Y227bXrbAm253t7//////D923dtmuwO23tb7//////yd22Z23a4btt77u3/////8B227tt3sAttjLt5//////EO9u626fENtve36f/////wZ52btuxgBtt29uP/////9AN79tt0AINtm17x//////ER7nbbdAQRtvt7Bf/////wBO37bZ4AAPbv7AP/////8mB7j22bLuQ7/nCJ//////AIf/326/8QH5AyAf/////xAsF/nuzTgngAMGf/////8GCAAPt3XuTgAJEB//////QJgAA7m3vzgHsYEf/////xARsADvvtmgIACIX/////8GUAAg9tt34ABAwB//////IBBBAt7t7sBAAJE//////wiSAAnvvb/pAASAH/////8AMAhJudexaSZTiJ//////JhlJJ9l87/iRv4A//////wCf9773796/3/yRH/////9IHr/gPrN3o/gIAF//////AQAAABnXbwAAiEg//////yBAAAAHYTgARokBP/////8IEgAAAMAAADAQIB//////AQCAAAAAAAAQZIif/////0BIAAAAAAAIDkAAP/////8QAgAIQAAAIAGJJj//////BiCIQQAEIQEAAACf///////////////////////8='" --stack-size=8kb .
