export class WebSocketService {
    private static socket: WebSocket;

    static connect = (hostUrl: string) => {
        this.socket = new WebSocket(`ws://${hostUrl}/ws/`);

        this.socket.onopen = () => {
            console.log(`WebSocket connection opened with -> ws://${hostUrl}/ws/`);
        };
        this.socket.onerror = (err) => {
            console.log("WebSocket error");
            console.log(err);
        }
    };
}