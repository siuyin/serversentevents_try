function listener() {
    const evSrc = new EventSource("events")
    const tm = document.getElementById("time")
    evSrc.onmessage = (ev) => {
        tm.textContent = `Current time on Server: ${ev.data}`
    }
}
listener()