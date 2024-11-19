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


