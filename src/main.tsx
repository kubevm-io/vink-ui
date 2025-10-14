import ReactDOM from 'react-dom/client'
import Console from '@/pages/console'
import '@/styles/main.css'

const params = new URLSearchParams(window.location.search)
const username = params.get('username') || 'hexiaodai'
const clientId = params.get('client_id') || '0'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Console username={username} clientId={clientId} />
)
