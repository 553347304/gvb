import {Message} from "@arco-design/web-vue";
import {Base} from "@/BAIYIN/HTTP";

function getLine(num: number): string {
    const err = new Error();
    const stack = err.stack?.split("\n");

    if (stack && stack.length > 2) {
        const callerLine: string = stack[num];
        const match = callerLine.match(/\(?(http[^)]+)\)?/);
        return match ? match[1] : "";
    }
    return "";
}

export function log(...s: any[]) {
    console.log(...s, getLine(3));
}

// if (!BaseLog(result)) return
export function BaseLog(result: Base.Response<any>): boolean {
    if (result.code !== 0) {
        Message.warning(result.message)
        return false
    }
    Message.success(result.message)
    return true
}
