export class UtilService {
    static delay = (ms: number) => new Promise(res => setTimeout(res, ms));
}