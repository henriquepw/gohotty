const queryString = window.location.search;
console.log({ queryString });
const urlParams = new URLSearchParams(queryString);
const ENDPOINT = urlParams.get("endpoint")?.toString() || "ws";

const loc = window.location;
let uri = loc.protocol === "https:" ? "wss:" : "ws:";
uri += `//${loc.host}`;
uri += `${loc.pathname}${ENDPOINT}`;

function tryConnect() {
	socket = new WebSocket(uri);
	socket.onopen = location.reload;
	socket.onerror = reconnect;
}

function reconnect() {
	setTimeout(tryConnect, 100);
}

let socket = new WebSocket(uri);
socket.onclose = reconnect;
