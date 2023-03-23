export class WebSocketService {
    private static _isConnected: boolean = false;
    private static _socket: WebSocket;

    private static _messageHandlers: { [name: string]: (data: any) => any };

    static connect = (originUrl: string, connectSuccess?: () => any) => {
        this._socket = new WebSocket(`${originUrl}/ws/`);

        this._socket.onopen = () => {
            this._isConnected = true;
            console.log(`WebSocket connection opened with -> ${originUrl}/ws/`);
            if (connectSuccess) {
                connectSuccess();
            }
        };
        this._socket.onerror = (event) => {
            this._isConnected = false;
            this._socket.close(1000, "Websocket error");
            console.log("WebSocket error");
            console.log(event);
        };
        this._socket.onclose = (event) => {
            this._isConnected = false;
            console.log(`Socket connection closed with code: ${event.code} and reason: ${event.reason}`);
        };
    };

    static emit = (event: string, data: any): boolean => {
        let success = false;

        if (this._isConnected) {
            try {
                this._socket.send(JSON.stringify({ event, data }));
                success = true;
            } catch (err) {
                console.log("Socket Emit Error:");
                console.log(err);
            }
        }

        return success;
    };
}