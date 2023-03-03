export const rupiahLocale = (x) => {
    return "Rp " + x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}

export const getDateNow = () => {
    return new Date().toLocaleString('en-us', {year: 'numeric', month: '2-digit', day: '2-digit'}).replace(/(\d+)\/(\d+)\/(\d+)/, '$3-$1-$2')
}