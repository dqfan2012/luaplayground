import './style.css';

import { RunLua } from '../wailsjs/go/app/App';

let scriptInput = document.getElementById('scriptInput');
let outputElement = document.getElementById('output');

window.runLua = function () {
    let script = scriptInput.value;

    if (script === "") return;

    RunLua(script)
        .then((result) => {
            console.log(result);
            outputElement.innerHTML = result;
        })
        .catch((err) => {
            console.error(err);
        });
};
