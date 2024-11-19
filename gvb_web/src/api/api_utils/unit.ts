import {nextTick} from "vue";


class _Unit {
    Scroll(selectors: string) {
        nextTick(() => {
            setTimeout(() => {
                let doc = document.querySelector(selectors)
                if (doc === null) return
                doc.scrollTo({
                    top: doc.scrollHeight,
                    behavior: "smooth", // 平滑效果
                })
            }, 0)
        })
    }
}

export const Unit = new _Unit()