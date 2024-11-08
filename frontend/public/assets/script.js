let term;
let socket;
let isConnected = false;
const PROXY_URL = 'ws://localhost:3001';

function initTerminal() {
    term = new Terminal({
        cursorBlink: true,
        theme: {
            background: '#000000',
            foreground: '#ffffff'
        },
        fontFamily: 'Menlo, Monaco, "Courier New", monospace',
        fontSize: 14,
        letterSpacing: 0,
        lineHeight: 1,
        scrollback: 1000,
        cols: 129,
        rows: 42
    });
    
    term.open(document.getElementById('terminal'));

    term.write('\r\x1B[1;32mTerminal Initialized Successfully!\x1B[0m\r\n\n');
    term.write('\r\x1B[37m# Hit the connect button above to run the following command\x1B[0m\r');
    term.write('\r\n\x1B[94m[visitor@sshfolio ~]$ \x1B[0mssh zachl.tech\n');

    const fitAddon = new FitAddon.FitAddon();
    term.loadAddon(fitAddon);
    
    fitAddon.fit();
    
    window.addEventListener('resize', () => {
        fitAddon.fit();
        if (socket && isConnected) {
            socket.send(JSON.stringify({
                type: 'resize',
                cols: term.cols,
                rows: term.rows
            }));
        }
    });
}

function updateButtons(connected) {
    document.getElementById('connectBtn').disabled = connected;
    document.getElementById('disconnectBtn').disabled = !connected;
}

function connect() {
    if (isConnected) return;
    
    updateStatus('Connecting...');
    
    try {
        socket = new WebSocket(PROXY_URL);
        
        socket.onopen = () => {
            isConnected = true;
            updateStatus('Connected');
            updateButtons(true);
            term.write('\r\n\x1B[1;32mConnected to zachl.tech\x1B[0m\r');
            
            term.onData(data => {
                if (isConnected) {
                    socket.send(data);
                }
            });
        };
        
        socket.onmessage = (event) => {
            term.write(event.data);
        };
        
        socket.onclose = () => {
            disconnect();
        };
        
        socket.onerror = (error) => {
            console.error('WebSocket error:', error);
            updateStatus('Connection error', 'error');
            disconnect();
        };
        
    } catch (error) {
        console.error('Connection error:', error);
        updateStatus('Failed to connect', 'error');
        disconnect();
    }
}

function disconnect() {
    if (!isConnected) return;
    
    if (socket) {
        socket.close();
    }
    
    isConnected = false;
    updateStatus('Disconnected');
    updateButtons(false);
    term.clear();
    term.write('\r\n\x1B[1;31mDisconnected from zachl.tech\x1B[0m\r\n\n');
    term.write('\r\x1B[37m# Hit the connect button above again to try it again :D\x1B[0m\r');
    term.write('\r\n\x1B[94m[visitor@sshfolio ~]$ \x1B[0mssh zachl.tech');
}

function updateStatus(message, type = 'info') {
    const statusElement = document.getElementById('status');
    statusElement.textContent = message;
    statusElement.style.color = type === 'error' ? '#ff4444' : '#ffffff';
}

window.addEventListener('load', () => {
    initTerminal();
});