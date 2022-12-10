export function isMobile(): boolean {
    const dv = document.getElementById('divscr');
    const sp = document.getElementById('res');

    return dv.offsetWidth - dv.clientWidth == 10
}