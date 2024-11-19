import {useStore} from "@/stores";
import {EventSourcePolyfill} from "event-source-polyfill";

const store = useStore();

class _SSE {

    private convert(data: any) {
        let params = ""
        Object.entries(data).forEach(([key, value]) => {
            params += `&${key}=${value}`
        });
        return params.slice(1)
    }

    get(url: string, params: any): EventSourcePolyfill {
        const eventSource = new EventSourcePolyfill(`${url}?${this.convert(params)}`, {
            heartbeatTimeout: 3 * 60 * 1000,
            headers: {
                token: store.userInfo.token,
                Accept: 'text/event-stream'
            },
            withCredentials: true,
        })
        eventSource.onerror = function (error) {
            // console.log(ev)
            eventSource.close()
        }
        return eventSource
    }



}

export const SSE = new _SSE()


// npm install @types/event-source-polyfill


// const eventSource = SSE.get("/api/big_model/chat_sse", chat)
// eventSource.onmessage = function (ev: MessageEvent) {
//     console.log(ev.data)
// }
// eventSource.onerror = function (ev: Event) {
//     console.log("传输完毕")
//     eventSource.close()
// }