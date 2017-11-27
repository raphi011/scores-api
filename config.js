const prod = process.env.NODE_ENV === 'production'

module.exports = {
  'process.env.BACKEND_URL': prod ? 'https://scores.raphi011.com' : 'http://localhost:3000'
}