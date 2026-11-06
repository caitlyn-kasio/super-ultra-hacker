package textImport

import (
	"math/rand"
)

func ReturnString() string {
	
	stringSlice := []string{}	

	stringSlice = append(stringSlice, "fn main() -> io::Result<()> {\n\tlet listener = TcpListener::bind(\"127.0.0.1:7878\")?;\n\tlet (tx, rx) = mpsc::channel();\n\tlet tx = Arc::new(Mutex::new(tx));\n\t\tfor (i, stream) in listener.incoming().enumerate() {\n\t\t\tmatch stream {\n\t\t\t\tOk(stream) => {\n\t\t\t\t\tlet tx = Arc::clone(&tx);\n\t\t\t\t\tthread::spawn(move || {\n\t\t\t\t\t\thandle_client(stream, tx.lock().unwrap().clone(), i);\n\t\t\t\t});\n\t\t\tOk(stream)\n\t\t}\n\t\tErr(e) => {\n\t\t\teprint!(\"Error accepting connection: {}\", e);\n\t\t}\n\t}\n\twhile let Ok(log_entry) = rx.recv() {\n\t\tprintln!(\"{}\", log_entry);\n\t}\n\tOk(())\n}")
	stringSlice = append(stringSlice, "use std::io::{self, Write};\nuse std::net::TcpListener;\nuse std::sync::mpsc;\nuse std::thread;\nuse std::time::SystemTime;\n\nfn handle_client(mut stream: std::net::TcpStream, sender: mpsc::Sender<String>) {\n\tlet mut buffer = [0; 512];\n\tloop {\n\t\tmatch stream.read(&mut buffer) {\n\t\t\tOk(n) if n == 0 => break,\n\t\t\tOk(n) => {\n\t\t\t\tlet data = String::from_utf8_lossy(&buffer[..n]);\n\t\t\t\tlet time = SystemTime::now()\n\t\t\t\t\t.duration_since(SystemTime::UNIX_EPOCH)\n\t\t\t\t\t.expect(\"Time went backwards\")\n\t\t\t\t\t.as_secs();\n\t\t\t\tlet log_entry = format!(\"Time: {}, Data: {}\n\", time, data);\n\t\t\t\tsender.send(log_entry.clone()).unwrap();\n\t\t\t\tstream.write_all(&data.as_bytes()).unwrap();\n\t\t\t}\n\t\tErr(e) => {\n\t\t\t\teprintln!(\"Error reading from stream: {}\", e);\n\t\t\t\tbreak;\n\t\t}\n\t}\n}\n\nfn main() -> io::Result<()> {\n\tlet listener = TcpListener::bind(\"127.0.0.1:7878\")?;\n\tlet (tx, rx) = mpsc::channel();\n\t\n\tfor stream in listener.incoming() {\n\t\tmatch stream {\n\t\t\tOk(stream) => {\n\t\t\t\tlet tx = tx.clone();\n\t\t\t\tthread::spawn(move || {\n\t\t\t\t\thandle_client(stream, tx);\n\t\t\t});\n\t\t\tOk(stream)\n\t\t}\n\t\tErr(e) => {\n\t\t\t\teprint!(\"Error accepting connection: {}\", e);\n\t\t}\n\t}\n\n\twhile let Ok(log_entry) = rx.recv() {\n\t\tprintln!(\"{}\", log_entry);\n\t}\n\n\tOk(())\n")
	stringSlice = append(stringSlice, )
	
	
	
	
	// Select a random slice
	sliceLen := len(stringSlice)
	
	randSlice := rand.Intn(int(sliceLen))
	
	return stringSlice[randSlice]
}



func ReturnChat() string {
	
	stringSlice := []string{}
	
	stringSlice = append(stringSlice, "hahah gskdhglsdjkgh ldsjghlajkh\n fskjhfgkas\nfkjdgfksdh\nflaksfblasjkfh\nfsahgfakshf")
	stringSlice = append(stringSlice, "neat giuksdhgjhs dgkndj lghd\nnew chat here fjsafgah")
	stringSlice = append(stringSlice, "wow I should really write new stuff\nfkgfaskhfg\naksjf\nlkdjglskdug\ngfhsdkg")
	
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