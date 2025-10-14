import { useEffect, useRef } from 'react'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import { WebLinksAddon } from 'xterm-addon-web-links'
import { WebglAddon } from 'xterm-addon-webgl'
import styles from '@/pages/console/index.module.less'
import 'xterm/css/xterm.css'

interface ConsoleProps {
    username: string
    clientId: string
}

const str2ab = (str: string) => {
    const enc = new TextEncoder()
    return enc.encode(str)
}

const Console: React.FC<ConsoleProps> = ({ username, clientId }) => {
    const host = window.location.host

    const divRef = useRef<HTMLDivElement>(null)
    const terminalRef = useRef<Terminal | null>(null)
    const wsRef = useRef<WebSocket | null>(null)

    const decoder = new TextDecoder('utf-8', { fatal: false })

    const initializeTerminal = () => {
        const terminal = new Terminal({
            cursorBlink: true,
            allowProposedApi: true,
            fontSize: 16,
            lineHeight: 1.2,
            theme: {
                foreground: '#F8F8F8',
                background: '#2D2E2C'
            },
            fontFamily: '"Cascadia Code", Menlo, monospace'
        })

        const fitAddon = new FitAddon()
        terminal.loadAddon(fitAddon)
        terminal.loadAddon(new WebLinksAddon())

        const ws = new WebSocket(`ws://${host}/ws/ssh?client_id=${clientId}&username=${username}`)
        ws.onopen = () => {
            terminal.open(divRef.current!)
            fitAddon.fit()
            terminal.loadAddon(new WebglAddon())
            terminal.focus()

            ws.send(str2ab("\n"))

            window.onresize = () => {
                fitAddon.fit()
            }
        }

        // FIXME: Need to retry
        ws.onerror = (err) => {
            console.log(err)
        }

        ws.onmessage = async (event) => {
            const rawData = await event.data.arrayBuffer()
            const text = decoder.decode(rawData, { stream: true })
            terminal.write(text)
        }

        ws.onerror = (err) => {
            terminal.write(`Websocket connection error: ${err}`)
        }

        terminal.onData((b) => {
            const msg = str2ab(b)
            ws.send(msg)
        })

        return { terminal, ws }
    }

    useEffect(() => {
        if (!clientId || !username) {
            return
        }

        const { terminal, ws } = initializeTerminal()

        terminalRef.current = terminal
        wsRef.current = ws

        return () => {
            ws.close()
            terminal.dispose()
            terminalRef.current = null
            wsRef.current = null
        }
    }, [clientId, username])

    return <div className={styles["console"]} ref={divRef} />
}

export default Console
