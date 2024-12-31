let term;
let socket;
let isConnected = false;
const PROXY_URL = 'wss://sshfolio-proxy.zachl.tech';

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
    });
    
    term.open(document.getElementById('terminal'));

    term.write('\r\x1B[1;32mTerminal Initialized Successfully!\x1B[0m\r\n\n');
    term.write('\r\x1B[37m# Hit the connect button above to run the following command\x1B[0m\r');
    term.write('\r\n\x1B[94m[visitor@sshfolio ~]$ \x1B[0mssh zachl.tech -p 2222');

    const fitAddon = new FitAddon.FitAddon();
    term.loadAddon(fitAddon);
    
    setTimeout(() => {
        fitAddon.fit();
        term.focus();
    }, 0);

    term.onKey(({ key, domEvent }) => {
        if (domEvent.ctrlKey && domEvent.key === 'c') {
            if (isConnected) {
                disconnect();
            }
        }
    });
    
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

function getWidth() {
    return Math.max(
        document.body.scrollWidth,
        document.documentElement.scrollWidth,
        document.body.offsetWidth,
        document.documentElement.offsetWidth,
        document.documentElement.clientWidth
    );
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
            term.write('\r\n\n\x1B[1;32mConnected to zachl.tech\x1B[0m\r');

            const dimensions = {
                cols: term.cols,
                rows: term.rows
            };

            socket.send(JSON.stringify({
                type: 'resize',
                cols: dimensions.cols,
                rows: dimensions.rows
            }));
            
            term.onData(data => {
                if (isConnected) {
                    if (data.includes("exit") || data.includes("\u0003")) {
                        disconnect();
                        return;
                    }
                    socket.send(JSON.stringify({
                        type: 'data',
                        data: data
                    }));
                }
            });
        };
        
        socket.onmessage = (event) => {
            try {
                const message = JSON.parse(event.data);
                if (message.type === 'status') {
                    switch(message.status) {
                        case 'disconnected':
                        case 'error':
                            isConnected = false;
                            disconnect();
                            return;
                    }
                }
                if (typeof event.data === 'string' && 
                    (event.data.includes("Connection closed") || 
                     event.data.includes("Session ended"))) {
                    disconnect();
                    return;
                }
                term.write(event.data);
            } catch (e) {
                term.write(event.data);
            }
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
    term.write('\r\n\x1B[94m[visitor@sshfolio ~]$ \x1B[0mssh zachl.tech -p 2222');
}

function updateStatus(message, type = 'info') {
    const statusElement = document.getElementById('status');
    statusElement.textContent = message;
    statusElement.style.color = type === 'error' ? '#ff4444' : '#ffffff';
}

window.addEventListener('load', () => {
    initTerminal();
});

async function copyButton() {
    const button = document.getElementById('clickToCopy');
    const textToCopy = 'ssh zachl.tech -p 2222';

    try {
        if (navigator.clipboard && navigator.clipboard.writeText) {
            await navigator.clipboard.writeText(textToCopy);
        } else {
            const textarea = document.createElement('textarea');
            textarea.value = textToCopy;
            textarea.style.position = 'fixed';
            textarea.style.opacity = '0';
            document.body.appendChild(textarea);
            textarea.select();
            document.execCommand('copy');
            document.body.removeChild(textarea);
        }
        
        button.textContent = 'Copied!';
    } catch (err) {
        console.error('Failed to copy: ', err);
        button.textContent = 'Copy failed';
    } finally {
        setTimeout(() => {
            button.textContent = 'Click to Copy!';
        }, 2000);
    }
}