import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime"
import "dayjs/locale/zh-cn.js"

dayjs.extend(relativeTime)
dayjs.locale("zh-cn")

export namespace TypeTime {
    export interface Time {
        Now: string
        NowDay: string
        Y: number
        M: number
        D: number
        h: number
        m: number
        s: number
        ms: number
    }

    export const Time = () => ({
        Now: "",
        NowDay: "",
        Y: 0,
        M: 0,
        D: 0,
        h: 0,
        m: 0,
        s: 0,
        ms: 0,
    })
}

class _Time {
    constructor() {
    }

    private s: number = 1000; // 毫秒
    private m: number = 60 * this.s; // 分钟
    private h: number = 60 * this.m; // 小时
    private D: number = 24 * this.h; // 天
    private M: number = 30 * this.D; // 天
    private Y: number = 365 * this.M; // 年

    Time(now: number | string = Date.now()): TypeTime.Time {
        // utc time convert
        if (typeof now === "string") {
            if (now.includes("+")) {
                const c: Date = new Date(now);
                now = c.getTime();
            }
        }
        const t: Date = new Date(now); // 创建 Date 对象
        const Y: number = t.getFullYear(); // 获取年份
        const M: number = t.getMonth() + 1; // 获取月份 (0-11)
        const D: number = t.getDate(); // 获取日期 (1-31)
        const h: number = t.getHours(); // 获取小时 (0-23)
        const m: number = t.getMinutes(); // 获取分钟 (0-59)
        const s: number = t.getSeconds(); // 获取秒 (0-59)
        const Now: string = `${Y}-${M}-${D}`
        const Day: string = `${h}:${m}:${s}`
        return {
            Now: `${Now} ${Day}`,
            NowDay: Day,
            Y: Y, M: M, D: D,
            h: h, m: m, s: s,
            ms: t.getTime(), // 时间戳
        };
    }

    Convert(ms: number): TypeTime.Time {
        const time = {
            Y: this.Y, M: this.M, D: this.D,
            h: this.h, m: this.m, s: this.s,
        };

        const result: any = {};
        for (const [t, v] of Object.entries(time)) {
            result[t] = Math.floor(ms / v);
            ms %= v;
        }
        const R = result as TypeTime.Time
        const Now: string = `${R.Y}-${R.M}-${R.D}`
        const Day: string = `${R.h}:${R.m}:${R.s}`
        return {
            Now: `${Now} ${Day}`,
            NowDay: Day,
            Y: R.Y, M: R.M, D: R.D,
            h: R.h, m: R.m, s: R.s,
            ms: ms, // 时间戳
        };
    }
}

export const Time = new _Time();
