class _Storage {
    constructor() {
    }

    Get(key: string) {
        return localStorage.getItem(key)
    }

    Set(key: string, v: string) {
        localStorage.setItem(key, v)
    }

    Remove(key: string) {
        localStorage.removeItem(key);
    }
}

export const Storage = new _Storage();