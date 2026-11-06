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
	
	stringSlice = append(stringSlice, "[red]error_404[white]: Wait, who are these people in the lobby?\n[red]coffee_mage[white]: They’re wearing matching sunglasses indoors.\n[red]error_404[white]: One of them just called the elevator “the escape pod.”\n[red]coffee_mage[white]: Maybe it’s a costume party?\n[red]error_404[white]: They’re asking which one of us is the “main character.”\n[red]coffee_mage[white]: I’m going to pretend I didn’t hear that.\n[red]error_404[white]: One is dramatically whispering into a disconnected phone.\n[red]coffee_mage[white]: I think they’re trying to look mysterious.\n[red]error_404[white]: They’ve put a sticky note on the coffee machine that says “ACCESS DENIED.”\n[red]coffee_mage[white]: That’s where we keep the decaf.\n[red]error_404[white]: Now someone’s asking if we have a “secure server room.”\n[red]coffee_mage[white]: We have a supply closet with a very loud fan.\n[red]error_404[white]: They’re calling it the core.\n[red]coffee_mage[white]: I’m calling facilities.\n[red]error_404[white]: One of them just gave the printer a codename.\n[red]coffee_mage[white]: Please tell me it’s not “The Beast.”\n[red]error_404[white]: It is absolutely “The Beast.”\n[red]coffee_mage[white]: I’m taking my lunch break in the other building.")
	stringSlice = append(stringSlice, "[blue]zero_cool[white]: The firewall just blinked. It knows we’re here.\n[blue]neon_ghost[white]: Relax. That was the screensaver.\n[blue]byte_bandit[white]: I’m seeing three layers of dramatic green text.\n[blue]zero_cool[white]: Make it four. The stakes demand it.\n[blue]neon_ghost[white]: Countdown at ten seconds!\n[blue]byte_bandit[white]: Ten seconds until what?\n[blue]neon_ghost[white]: No idea, but the music is getting louder.\n[blue]zero_cool[white]: We’re in. The vending machine is dispensing snacks.\n[blue]byte_bandit[white]: Mission accomplished. Get the chips.")
	stringSlice = append(stringSlice, "[blue]zero_cool[white]: The terminal is glowing. This is either progress or a very expensive screensaver.\n[blue]neon_ghost[white]: I have initiated the ultra-secret dramatic typing sequence.\n[blue]byte_bandit[white]: You’re typing “password123” into a calculator.\n[blue]neon_ghost[white]: It’s called misdirection.\n[blue]zero_cool[white]: The big red button is blinking.\n[blue]byte_bandit[white]: It says “Do Not Press.”\n[blue]neon_ghost[white]: Exactly. They want us to think it’s a button.\n[blue]zero_cool[white]: It is very clearly a button.\n[blue]byte_bandit[white]: The countdown has appeared on every monitor.\n[blue]neon_ghost[white]: What does it count down to?\n[blue]byte_bandit[white]: The pizza delivery. It’s two minutes away.\n[blue]zero_cool[white]: Stay focused. The fate of the mission is at stake.\n[blue]neon_ghost[white]: I’m seeing a mysterious blinking cursor.\n[blue]byte_bandit[white]: That means the computer is waiting for you to type something.\n[blue]neon_ghost[white]: I knew it. It’s a staring contest.\n[blue]zero_cool[white]: We’re in! The screen says “Welcome, Admin.”\n[blue]byte_bandit[white]: That’s the login screen at the library.\n[blue]neon_ghost[white]: Mission accomplished. Let’s get out before the printer jams.")
	stringSlice = append(stringSlice, "[red]packet_panic[white]: Who are the people in matching jackets by the arcade cabinet?\n[red]syntax_sorcerer[white]: The ones arguing with the claw machine?\n[red]packet_panic[white]: They said they’re our rivals.\n[red]syntax_sorcerer[white]: Their leader just introduced himself as “Lord Lag.”\n[red]packet_panic[white]: He challenged us to a duel.\n[red]syntax_sorcerer[white]: At the foosball table. Very high stakes.\n[red]packet_panic[white]: They brought a laptop with a sticker that says “I void warranties.”\n[red]syntax_sorcerer[white]: It’s plugged into the wall but the screen is off.\n[red]packet_panic[white]: One of them is typing furiously anyway.\n[red]syntax_sorcerer[white]: Maybe they’re composing a very intense email.\n[red]packet_panic[white]: Lord Lag says he’s “breached the snack perimeter.”\n[red]syntax_sorcerer[white]: He found the bowl of mints.\n[red]packet_panic[white]: They’re demanding we reveal our secret code names.\n[red]syntax_sorcerer[white]: Tell them mine is “The Person Who Knows Where the Exit Is.”\n[red]packet_panic[white]: They’re powering up a fog machine.\n[red]syntax_sorcerer[white]: That’s the meeting room humidifier.\n[red]packet_panic[white]: I think the duel is starting.\n[red]syntax_sorcerer[white]: Great. I call blue team.")
	
	sliceLen := len(stringSlice)
	randSlice := rand.Intn(int(sliceLen))
	
	return stringSlice[randSlice]
}