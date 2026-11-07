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


	
	stringSlice = append(stringSlice, "IRC Viewer 12.41.12\nJoined Channel Sys7emSh0ck\n[blue::b]z0mg_z0rg[white::B]:Hey, we're getting a lot of network scans from all over. Think it's an attack?\n[blue::b]N00bSlayer99[white::B]:think its cold boot???? maybe they wanna get us back\n[blue::b]z0mg_z0rg[white::B]:What did you do?\n[blue::b]N00bSlayer99[white::B]:Redirected their pizza delivery\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Have they found a way in yet?\n[blue::b]z0mg_z0rg[white::B]:The IDS hasnt picked up anything, I'm looking now.\n[blue::b]leetKev[white::B]:Hey did one of you hack the tv? why is it showing a snowflake\n[blue::b]N00bSlayer99[white::B]:oh shit\n[blue::b]xX_Shad0wByt3_Xx[white::B]:Did they get to the server?\n[blue::b]N00bSlayer99[white::B]:should we counter-hack them???\n[blue::b]z0mg_z0rg[white::B]:Theres an open connection to the database, is that them?\n[blue::b]leetKev[white::B]:Shut the server down!\nConnection has been lost...\nEND")
	stringSlice = append(stringSlice, "IRC Viewer 12.41.12\nJoined Channel Sys7emSh0ck\n[blue::b]leetKev[white::B]:Why is the office printer making that noise?\n[blue::b]z0mg_z0rg[white::B]:What noise?\n[blue::b]leetKev[white::B]:It just printed a photo of a duck wearing sunglasses.\n[blue::b]N00bSlayer99[white::B]:thats our new security mascot\n[blue::b]xX_Shad0wByt3_Xx[white::B]:I thought the mascot was a raccoon.\n[blue::b]N00bSlayer99[white::B]:the raccoon retired\n[blue::b]z0mg_z0rg[white::B]:I’m checking the printer logs. Everything looks normal.\n[blue::b]leetKev[white::B]:The duck has a name tag. It says “Gary.”\n[blue::b]xX_Shad0wByt3_Xx[white::B]:We have a Gary?\n[blue::b]z0mg_z0rg[white::B]:Not that I know of.\n[blue::b]N00bSlayer99[white::B]:maybe Gary is the one scanning us\n[blue::b]leetKev[white::B]:There’s another page coming out.\n[blue::b]xX_Shad0wByt3_Xx[white::B]:What does it say?\n[blue::b]leetKev[white::B]:“Please refill the paper tray.”\n[blue::b]z0mg_z0rg[white::B]:That’s just the printer.\n[blue::b]N00bSlayer99[white::B]:or is it\n[blue::b]xX_Shad0wByt3_Xx[white::B]:I’m naming the duck Gary either way.\n[blue::b]leetKev[white::B]:Gary just printed his own badge.\n[blue::b]z0mg_z0rg[white::B]:Someone please unplug the printer.\nConnection has been lost...")
	stringSlice = append(stringSlice, "[red]i[blue]r[green]c[red]s[blue]n[green]0[red]0[blue]p[green]e[red]r\ncopyright xX_Shad0wByt3_Xx[white]\nconnecting to channel cold boot\nbreaching...\nsuccess!\n[red::b]S1mp1e[white::B]:you found a way into beacon??\n[red::b]CupCakeCrasher[white::B]:omg no way how\n[red::b]XxX_R3KT_XxX[white::B]:It was easy... they left an old login endpoint open and didn't update it.\n[red::b]Sn3akyP3tty[white::B]:So what's the plan now?\n[red::b]XxX_R3KT_XxX[white::B]:We have to make sure to get in and get out before systemshock finds it.\n[red::b]S1mp1e[white::B]:as if they could\n[red::b]CupCakeCrasher[white::B]:did you see the priv esc that dropped last week? maybe they still havent updated for that\n[red::b]S1mp1e[white::B]:youre right, im in.\n[red::b]S1mp1e[white::B]:shit. somethings wrong. it knows were here\n[red::b]XxX_R3KT_XxX[white::B]:Don't touch anything. Let it pass.\n[red::b]Sn3akyP3tty[white::B]:Did they finally wake up?\n[red::b]CupCakeCrasher[white::B]:i left them a cupcake :3\n[red::b]Sn3akyP3tty[white::B]:Our logo is a snowflake your're supposed to leave a snowflake\n[red::b]CupCakeCrasher[white::B]:whoops\n[red::b]S1mp1e[white::B]:im through to the print queue, can steal anything they try to print\n[red::b]Server:WE SEE YOU\n[white::B]Connection failed.\nEND")
	stringSlice = append(stringSlice, "[green::b]Welcome to BeaconChat!\nConnecting to channel OfficeOps...\n[green::b]j_dobbs[white::B]:Morning, everyone. The dashboard is showing a strange alert. Mike, can you take a look?\n[green::b]mike_t[white::B]:Already on it. The system’s responding normally.\n[green::b]sandra_k[white::B]:Then why did three invoices just get marked “reviewed”?\n[green::b]pamela_g[white::B]:I thought Joan reviewed them. She’s very thorough.\n[green::b]j_dobbs[white::B]:I haven’t seen those invoices.\n[green::b]mike_t[white::B]:Okay. That’s not normal. I’m checking the activity log.\n[green::b]sandra_k[white::B]:There’s another discrepancy. Someone opened the old vendor folder.\n[green::b]pamela_g[white::B]:Maybe it was the printer again?\n[green::b]mike_t[white::B]:Printers don’t open folders, Pamela.\n[green::b]pamela_g[white::B]:This one has been acting very confidently lately.\n[green::b]j_dobbs[white::B]:Let’s keep this calm. Mike, pause access to that folder while you investigate.\n[green::b]mike_t[white::B]:Done. The alert just cleared.\n[green::b]sandra_k[white::B]:That was fast.\n[green::b]mike_t[white::B]:Too fast. Someone else cleared it.\n[green::b]j_dobbs[white::B]:Who has access to the alert panel?\n[green::b]benny_f[white::B]:I was just checking the logs. Everything’s fine now.\n[green::b]mike_t[white::B]:You shouldn’t have been able to clear that alert.\n[green::b]benny_f[white::B]:Maybe the system knows when it’s being watched.\nConnection Closed...\nEND")
	
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