// server.js
const WebSocket = require('ws');
const { Client } = require('ssh2');
const http = require('http');

const server = http.createServer();
const wss = new WebSocket.Server({ server });

wss.on('connection', (ws) => {
    console.log('New connection established');
    
    const ssh = new Client();
    let stream = null;
    
    ssh.on('ready', () => {
        console.log('SSH Connection established');
        
        ssh.shell({ 
            term: 'xterm-256color',
            cols: 80,
            rows: 24
        }, (err, _stream) => {
            if (err) {
                console.error('Shell error:', err);
                ws.close();
                return;
            }
            
            stream = _stream;
            
            // Send SSH data to WebSocket client
            stream.on('data', (data) => {
                try {
                    ws.send(data.toString('utf8'));
                } catch (error) {
                    console.error('WebSocket send error:', error);
                }
            });
            
            // Send WebSocket data to SSH connection
            ws.on('message', (data) => {
                try {
                    // Check if the message is a resize command
                    const message = data.toString();
                    if (message.startsWith('{')) {
                        const parsed = JSON.parse(message);
                        if (parsed.type === 'resize') {
                            stream.setWindow(parsed.rows, parsed.cols);
                            return;
                        }
                    }
                    // Otherwise treat it as regular input
                    stream.write(data);
                } catch (error) {
                    console.error('Error handling message:', error);
                }
            });
            
            // Handle stream close
            stream.on('close', () => {
                console.log('Stream closed');
                ssh.end();
                ws.close();
            });
        });
    });
    
    ssh.on('error', (err) => {
        console.error('SSH error:', err);
        ws.close();
    });
    
    ssh.connect({
        host: 'zachl.tech',
        username: 'anonymous',
        tryKeyboard: true,
        password: ''
    });
    
    ws.on('close', () => {
        console.log('WebSocket connection closed');
        if (stream) stream.close();
        ssh.end();
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

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
    console.log(`WebSocket server listening on port ${PORT}`);
});