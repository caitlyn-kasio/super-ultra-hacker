package textImport

import (
	"math/rand"
)

func ReturnString() string {
	
	stringSlice := []string{}	

	stringSlice = append(stringSlice, "fn main() -> io::Result<()> {\n\tlet listener = TcpListener::bind(\"127.0.0.1:7878\")?;\n\tlet (tx, rx) = mpsc::channel();\n\tlet tx = Arc::new(Mutex::new(tx));\n\t\tfor (i, stream) in listener.incoming().enumerate() {\n\t\t\tmatch stream {\n\t\t\t\tOk(stream) => {\n\t\t\t\t\tlet tx = Arc::clone(&tx);\n\t\t\t\t\tthread::spawn(move || {\n\t\t\t\t\t\thandle_client(stream, tx.lock().unwrap().clone(), i);\n\t\t\t\t});\n\t\t\tOk(stream)\n\t\t}\n\t\tErr(e) => {\n\t\t\teprint!(\"Error accepting connection: {}\", e);\n\t\t}\n\t}\n\twhile let Ok(log_entry) = rx.recv() {\n\t\tprintln!(\"{}\", log_entry);\n\t}\n\tOk(())\n}")
	stringSlice = append(stringSlice, "use std::io::{self, Write};\nuse std::net::TcpListener;\nuse std::sync::mpsc;\nuse std::thread;\nuse std::time::SystemTime;\n\nfn handle_client(mut stream: std::net::TcpStream, sender: mpsc::Sender<String>) {\n\tlet mut buffer = [0; 512];\n\tloop {\n\t\tmatch stream.read(&mut buffer) {\n\t\t\tOk(n) if n == 0 => break,\n\t\t\tOk(n) => {\n\t\t\t\tlet data = String::from_utf8_lossy(&buffer[..n]);\n\t\t\t\tlet time = SystemTime::now()\n\t\t\t\t\t.duration_since(SystemTime::UNIX_EPOCH)\n\t\t\t\t\t.expect(\"Time went backwards\")\n\t\t\t\t\t.as_secs();\n\t\t\t\tlet log_entry = format!(\"Time: {}, Data: {}\n\", time, data);\n\t\t\t\tsender.send(log_entry.clone()).unwrap();\n\t\t\t\tstream.write_all(&data.as_bytes()).unwrap();\n\t\t\t}\n\t\tErr(e) => {\n\t\t\t\teprintln!(\"Error reading from stream: {}\", e);\n\t\t\t\tbreak;\n\t\t}\n\t}\n}\n\nfn main() -> io::Result<()> {\n\tlet listener = TcpListener::bind(\"127.0.0.1:7878\")?;\n\tlet (tx, rx) = mpsc::channel();\n\t\n\tfor stream in listener.incoming() {\n\t\tmatch stream {\n\t\t\tOk(stream) => {\n\t\t\t\tlet tx = tx.clone();\n\t\t\t\tthread::spawn(move || {\n\t\t\t\t\thandle_client(stream, tx);\n\t\t\t});\n\t\t\tOk(stream)\n\t\t}\n\t\tErr(e) => {\n\t\t\t\teprint!(\"Error accepting connection: {}\", e);\n\t\t}\n\t}\n\n\twhile let Ok(log_entry) = rx.recv() {\n\t\tprintln!(\"{}\", log_entry);\n\t}\n\n\tOk(())\n")
	stringSlice = append(stringSlice, "")
	
	
	
	
	// Select a random slice
	sliceLen := len(stringSlice)
	
	randSlice := rand.Intn(int(sliceLen))
	
	return stringSlice[randSlice]
}



func ReturnChat() string {
	
	stringSlice := []string{}

	
	stringSlice = append(stringSlice, "[red]i[blue]r[green]c[red]s[blue]n[green]0[red]0[blue]p[green]e[red]r\ncopyright xX_Shad0wByt3_Xx[white]\nconnecting to channel cold boot\nbreaching...\nsuccess!\n[red::b]S1mp1e[white::B]:you found a way into beacon??\n[red::b]CupCakeCrasher[white::B]:omg no way how\n[red::b]XxX_R3KT_XxX[white::B]:It was easy... they left an old login endpoint open and didn't update it.\n[red::b]Sn3akyP3tty[white::B]:So what's the plan now?\n[red::b]XxX_R3KT_XxX[white::B]:We have to make sure to get in and get out before systemshock finds it.\n[red::b]S1mp1e[white::B]:as if they could\n[red::b]CupCakeCrasher[white::B]:did you see the priv esc that dropped last week? maybe they still havent updated for that\n[red::b]S1mp1e[white::B]:youre right, im in.\n[red::b]S1mp1e[white::B]:shit. somethings wrong. it knows were here\n[red::b]XxX_R3KT_XxX[white::B]:Don't touch anything. Let it pass.\n[red::b]Sn3akyP3tty[white::B]:Did they finally wake up?\n[red::b]CupCakeCrasher[white::B]:i left them a cupcake :3\n[red::b]Sn3akyP3tty[white::B]:Our logo is a snowflake your're supposed to leave a snowflake\n[red::b]CupCakeCrasher[white::B]:whoops\n[red::b]S1mp1e[white::B]:im through to the print queue, seeing files talking about something called Project Gamma??\n[red::b]Server:WE SEE YOU\n[white::B]Connection failed.\nEND")
	stringSlice = append(stringSlice, "[green::b]Welcome to BeaconChat!\nConnecting to channel OfficeOps...\n[green::b]joan_d[white::B]:Morning, everyone. The dashboard is showing a strange alert. Mike, can you take a look?\n[green::b]mike_t[white::B]:Already on it. The system’s responding normally.\n[green::b]sandra_k[white::B]:Then why did three invoices just get marked “reviewed”?\n[green::b]pamela_g[white::B]:I thought Joan reviewed them. She’s very thorough.\n[green::b]joan_d[white::B]:I haven’t seen those invoices.\n[green::b]mike_t[white::B]:Okay. That’s not normal. I’m checking the activity log.\n[green::b]sandra_k[white::B]:There’s another discrepancy. Someone opened the old vendor folder.\n[green::b]pamela_g[white::B]:Maybe it was the printer again?\n[green::b]mike_t[white::B]:Printers don’t open folders, Pamela.\n[green::b]pamela_g[white::B]:This one has been acting very confidently lately.\n[green::b]joan_d[white::B]:Let’s keep this calm. Mike, pause access to that folder while you investigate.\n[green::b]mike_t[white::B]:Done. The alert just cleared.\n[green::b]sandra_k[white::B]:That was fast.\n[green::b]mike_t[white::B]:Too fast. Someone else cleared it.\n[green::b]joan_d[white::B]:Who has access to the alert panel?\n[green::b]benny_f[white::B]:I was just checking the logs. Everything’s fine now.\n[green::b]mike_t[white::B]:You shouldn’t have been able to clear that alert.\n[green::b]benny_f[white::B]:Maybe the system knows when it’s being watched.\nConnection Closed...\nEND")
	stringSlice = append(stringSlice, "IRC Viewer 12.41.12\nJoined Channel Sys7emSh0ck\n[blue::b]leetKev[white::B]:Why is there music coming from the server room\n[blue::b]leetKev[white::B]:Its really ominous im kind of scared\n[blue::b]N00bSlayer99[white::B]:think its cold boot?\n[blue::b]xX_Shad0wByt3_Xx[white::B]:I hear it too. Concerning.\n[blue::b]z0mg_z0rg[white::B]:Oh yeah thats my old computer!\n[blue::b]N00bSlayer99[white::B]:what\n[blue::b]z0mg_z0rg[white::B]:Yeah remember the old old server? It woke up today and started playing music!\n[blue::b]N00bSlayer99[white::B]:dead servers shouldnt play carnival music\n[blue::b]leetKev[white::B]:thats terrifying shut it down\n[blue::b]z0mg_z0rg[white::B]:I tried but it just rebooted, and plus its extra processing power so Im not worried\n[blue::b]xX_Shad0wByt3_Xx[white::B]:It's 4 gigs of RAM, tops. Just unplug it?\n[blue::b]N00bSlayer99[white::B]:no keep it!!!! maybe it can help us\n[blue::b]N00bSlayer99[white::B]:maybe it knows things.........\n[blue::b]xX_Shad0wByt3_Xx[white::B]:It's a security liability and the music is, frankly, deeply unsettling.\n[blue::b]leetKev[white::B]:Yeah it sounds like its trying to tell us something I dont like it\n[blue::b]z0mg_z0rg[white::B]:It just printed a system message saying Im back\n[blue::b]leetKev[white::B]:?????? KILL IT\n[blue::b]z0mg_z0rg[white::B]:It comes back to life and you wanna kill it?\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Kill it.\n[blue::b]z0mg_z0rg[white::B]:Fine.\nConnection Closed...\nEND")
	stringSlice = append(stringSlice, "[green::b]Welcome to BeaconChat!\nConnecting to channel OfficeOps...\n[green::b]sandra_k[white::B]:Hey, I just found a file called 'DO NOT FINISH', anyone know what's up with that?\n[green::b]pamela_g[white::B]:Oh yeah that thing! That was fun.\n[green::b]sandra_k[white::B]:What was it?\n[green::b]mike_t[white::B]:Oh yeah project Omega! That *was* fun.\n[green::b]pamela_g[white::B]:Good times. I miss Darrel he was so sweet.\n[green::b]sandra_k[white::B]:Im sorry but what was project omega? I opened the file and it's just a big spreadsheet.\n[green::b]joan_d[white::B]:Oh yeah Darrel! He was a lot of fun, hope he's doing ok.\n[green::b]pamela_g[white::B]:Hey Joan why did we stop working on that again? We were really close!\n[green::b]joan_d[white::B]:Something about IT security?\n[green::b]mike_t[white::B]:I never sent anything about omega? What did it say?\n[green::b]sandra_k[white::B]:I still have no idea what you're all talking about!\n[green::b]benny_f[white::B]:It was closed for good reason. Very insecure file.\n[green::b]joan_d[white::B]:Benny you can't just send messages as IT.\n[green::b]benny_f[white::B]:I didn't.\n[green::b]pamela_g[white::B]:Then who closed it???\n[green::b]sandra_k[white::B]:Hey why did someone draw a lightning bolt across like half of this thing?\n[green::b]joan_d[white::B]:We probably shouldn't keep that around. Mike, should we delete that?\n[green::b]mike_t[white::B]:We shouldn't for policy reasons. But...\nConnection Closed...")
	stringSlice = append(stringSlice, "IRC Viewer 12.41.12\nJoined Channel Sys7emSh0ck\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Just found a way into BeaconCorp.\n[blue::b]N00bSlayer99[white::B]:omg what how\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Phished some credentials, looks like we have a login as 'Pamela G''\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Her password was password12345.\n[blue::b]z0mg_z0rg[white::B]:Damn, should've just tried to brute force them.\n[blue::b]leetKev[white::B]:just tried it it works!\n[blue::b]N00bSlayer99[white::B]:go slow! we dont wanna wake the system up before we take a look around\n[blue::b]z0mg_z0rg[white::B]:Just got the server listening, send a connection so we can get a reverse shell\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Already on it...\n[blue::b]leetKev[white::B]:just found something called project omega i knew they were up to something\n[blue::b]N00bSlayer99[white::B]:kev we already pwned that i dont think they even noticed\n[blue::b]z0mg_z0rg[white::B]:Cool, looks like were in. Im scanning for any escalations\n[blue::b]z0mg_z0rg[white::B]:Wait, someone just connected through the rshell??\n[blue::b]xX_Shad0wByt3_Xx[white::B]:That can't be possible, are you sure?\n[blue::b]N00bSlayer99[white::B]:cold boot...\nServer: Server is shutting down in 10 seconds...\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Damn, so close...\nConnection Closed...\nEND")
	
	sliceLen := len(stringSlice)
	randSlice := rand.Intn(int(sliceLen))
	
	return stringSlice[randSlice]
}

func ReturnCandy(sliceint int) string {
	
	stringSlice := []string{}
	
	stringSlice = append(stringSlice, "[red::b]Breaching the mainframe")
	stringSlice = append(stringSlice, "[red::b]Pwning n00bs")
	stringSlice = append(stringSlice, "[red::b]Gaining access")
	stringSlice = append(stringSlice, "[red::b]Executing exploit")
	stringSlice = append(stringSlice, "[red::b]Finding zero-day")
	stringSlice = append(stringSlice, "[red::b]Initiating protocol [green::b]Omega[red::b]")
	stringSlice = append(stringSlice, "[red::b]Subverting firewall")
	stringSlice = append(stringSlice, "[red::b]Overwriting memory")
	stringSlice = append(stringSlice, "[red::b]Reading passwords")
	stringSlice = append(stringSlice, "[red::b]Injecting code")
	stringSlice = append(stringSlice, "[red::b]Escalating priveleges")
	stringSlice = append(stringSlice, "[red::b]Hijacking session")
	stringSlice = append(stringSlice, "[red::b]Nullifying IDS")
	stringSlice = append(stringSlice, "[red::b]Traversing air-gap")
	stringSlice = append(stringSlice, "[red::b]Decrypting cypher")
	
		//Template	
// 		stringSlice = append(stringSlice, "[red::b]")	

	
	return stringSlice[sliceint]
}