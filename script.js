
const form = document.getElementById('form');
const output = document.getElementById('output')
form.addEventListener('submit', function(event) {
    const weight = parseFloat(event.target.elements[1].value)
    const bar = parseInt(event.target.elements[0].value)
    const plates = GetPlatesForBarJS(weight, bar)
    
    while (output.hasChildNodes()) {
        output.firstChild.remove()
    }
    output.innerText = 'Plates: ' + plates.join(", ")

    for (var plate of plates.reverse()) {
        const el = document.createElement('div')
        el.classList.add('plate')
        el.classList.add(plate >= 5 ? 'bumper' : 'training')
        el.classList.add(('kg' + plate).replace(".", "-"))
        output.appendChild(el)
    }

    event.preventDefault()
});
