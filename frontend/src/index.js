import './style.css';

import { RunLua } from '../wailsjs/go/app/App';

let scriptInput = document.getElementById('scriptInput');
let outputElement = document.getElementById('output');

// console.log(document.getElementById('runButton'));
// console.log(document.getElementById('scriptInput'));
// console.log(document.getElementById('output'));

window.runLua = function () {
    let script = scriptInput.value;

    if (script === "") return;

    try {
        RunLua(script)
            .then((result) => {
                outputElement.innerHTML = result;
            })
            .catch((err) => {
                console.error(err);
            })
    } catch(err) {
        console.log(err);
    }
};
