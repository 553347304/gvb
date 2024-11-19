import type {Plugin} from 'postcss'

export interface Options {
    viewPortWidth?: number;
    mediaQuery?: boolean;
    unitToConvert?: string;
}

const defaultOptions = {
    viewPortWidth: 375,
    mediaQuery: false,
    unitToConvert: 'px'
}

let reg = /(\d+\.?\d?)px/g;
export const pxtoViewport = (options: Options = defaultOptions): Plugin => {
    const opt = Object.assign({}, defaultOptions, options)
    return {
        postcssPlugin: 'postcss-px-to-viewport',
        //css节点都会经过这个钩子
        Declaration(node: any) {
            const value = node.value

            //匹配到px 转换成vw
            // if (node.source.input.file.includes("node_modules")){
            //     return
            // }
            if (value.includes(opt.unitToConvert)) {
                let val = value.replace(reg, (matched: any, capture1: string) => {
                    const transformValue = (Number(capture1) / opt.viewPortWidth) * 100
                    return `${transformValue.toFixed(2)}vw`;
                })
                node.value = val
            }
        },
    }
}


