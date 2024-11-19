export const Type = new class _Type {
    ListString(list: []): string {
        let str: string = "";
        for (let i = 0; i < list.length; i++) {
            str += list[i] + "\n";
        }
        return str
    }

    StringList(s: string, char: string): string[] {
        return s.split(char)
    }
}


/*
string.charAt(str.length - 1);  // 获取头部第一个字符串
string.charAt(0);               // 获取尾部第一个字符串
string.slice(1);                // 删除头部第一个字符串
string.slice(0, -1);            // 删除尾部第一个字符



 */




