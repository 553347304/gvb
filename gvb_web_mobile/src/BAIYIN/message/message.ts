import {showNotify} from 'vant';

class _Message {
    primary = (message: string) => showNotify({type: 'primary', message: message});
    success = (message: string) => showNotify({type: 'success', message: message});
    warning = (message: string) => showNotify({type: 'warning', message: message});
    error = (message: string) => showNotify({type: 'danger', message: message});
}

export const Message = new _Message()
