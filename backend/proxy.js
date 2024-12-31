const WebSocket = require('ws');
const { Client } = require('ssh2');
const http = require('http');

const server = http.createServer();
const wss = new WebSocket.Server({ server });

wss.on('connection', (ws) => {
    console.log('New connection established');
    
    const ssh = new Client();
    let stream = null;
    let initialDimensionsReceived = false;
    let terminalDimensions = { cols: 80, rows: 24 };

    function cleanupConnection() {
        try {
            if (stream) {
                stream.removeAllListeners();
                stream.end();
            }
            if (ssh) {
                ssh.removeAllListeners();
                ssh.end();
            }
        } catch (error) {
            console.error('Cleanup error:', error);
        }
    }

    ws.on('message', (data) => {
        try {
            const message = data.toString();
            if (message.startsWith('{')) {
                const parsed = JSON.parse(message);
                if (parsed.type === 'resize') {
                    terminalDimensions.cols = parsed.cols;
                    terminalDimensions.rows = parsed.rows / 1.1;
                    if (!initialDimensionsReceived) {
                        initialDimensionsReceived = true;
                        initializeSSHConnection();
                    } else if (stream) {
                        stream.setWindow(parsed.rows, parsed.cols);
                    }
                    return;
                } else if (parsed.type === 'data' && stream) {
                    stream.write(parsed.data);
                    return;
                }
            }
            if (stream) {
                stream.write(data);
            }
        } catch (error) {
            console.error('Message handling error:', error);
        }
    });

    function initializeSSHConnection() {
        ssh.connect({
            host: 'zachl.tech',
            username: 'visitor',
            port: 2222,
            tryKeyboard: true,
            password: ''
        });
    }
    
    ssh.on('ready', () => {
        ws.send(JSON.stringify({ type: 'status', status: 'connected' }));
        console.log('SSH Connection established');
        
        ssh.shell({ 
            term: 'xterm-256color',
            cols: terminalDimensions.cols,
            rows: terminalDimensions.rows
        }, (err, sstream) => {
            if (err) {
                console.error('Shell error:', err);
                ws.send(JSON.stringify({ type: 'status', status: 'error', message: err.message }));
                return;
            }
            
            stream = sstream;
            
            stream.on('data', (data) => {
                if (ws.readyState === ws.OPEN) {
                    ws.send(data.toString('utf8'));
                }
            });
            
            stream.on('close', () => {
                console.log('Stream closed');
                if (ws.readyState === ws.OPEN) {
                    ws.send(JSON.stringify({ type: 'status', status: 'disconnected' }));
                }
                cleanupConnection();
            });

            stream.on('end', () => {
                console.log('Stream ended');
                if (ws.readyState === ws.OPEN) {
                    ws.send(JSON.stringify({ type: 'status', status: 'disconnected' }));
                }
                cleanupConnection();
            });
        });
    });
    
    ssh.on('error', (err) => {
        console.error('SSH error:', err);
        if (ws.readyState === ws.OPEN) {
            ws.send(JSON.stringify({ type: 'status', status: 'error', message: err.message }));
        }
        cleanupConnection();
    });
    
    ssh.connect({
        host: 'zachl.tech',
        port: 2222,
        username: 'visitor',
        tryKeyboard: true,
        password: ''
    });
    
    wss.on('close', () => {
        console.log('WebSocket connection closed');
        cleanupConnection();
    });
});

server.on('upgrade', (request, socket, head) => {
    const headers = {
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'GET, POST',
        'Access-Control-Allow-Headers': 'Content-Type'
    };

    if (request.method === 'OPTIONS') {
        socket.write('HTTP/1.1 204 No Content\r\n');
        Object.entries(headers).forEach(([key, value]) => {
            socket.write(`${key}: ${value}\r\n`);
        });
        socket.write('\r\n');
        socket.destroy();
        return;
    }
});

const PORT = 3001;
server.listen(PORT, () => {
    console.log(`WebSocket server listening on port ${PORT}`);
});

server.on('close', () => {
    wss.clients.forEach(client => {
        if (client.readyState === WebSocket.OPEN) {
            client.close();
        }
    });
});