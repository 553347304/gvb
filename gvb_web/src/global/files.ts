class _File {
    head: string = "https://vip.123pan.cn/1821560246/website/image/head/koneko.png";

    icon: { [key: string]: string } = {
        bilibili: "https://vip.123pan.cn/1821560246/website/image/app/bilibili.png",
        wx: "https://vip.123pan.cn/1821560246/website/image/app/wx.png",
        qq: "https://vip.123pan.cn/1821560246/website/image/app/qq.png",
        githup: "https://vip.123pan.cn/1821560246/website/image/app/githup.png",
        gitee: "https://vip.123pan.cn/1821560246/website/image/app/gitee.png",
        filings: "https://vip.123pan.cn/1821560246/website/image/app/filings.png",
    };
    bg: string[] = [
        "https://vip.123pan.cn/1821560246/website/image/bg/Konachan_19.jpg",
        "https://vip.123pan.cn/1821560246/website/image/bg/Konachan_25.jpg",
    ];
}

export const Files = new _File()