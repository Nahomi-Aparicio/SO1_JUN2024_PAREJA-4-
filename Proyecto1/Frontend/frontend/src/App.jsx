import React, { useState, useEffect } from 'react';

import { createProcess,deleteProcess} from "./api/Endpoint.jsx";
import PASTEL from './utualito.jsx'; // Importación correcta
import TablaProcesos from './tabla.jsx'; // Importación correcta
import './estilos_pagina.css'; 
import TablaPid from './TablaPid';
function App() {
  const [actionButton, setActionButton] = useState("Crear Proceso");
  const [pid, setPid] = useState(0);
  const [entradaTexto, setEntradaTexto] = useState('');
  const [procesos, setProcesos] = useState([]);


const createProc = async () => {
  try {
      const req = await createProcess();
      const res = await req.json();
      setPid(res.pid);
      setProcesos((prevProcesos) => [
        ...prevProcesos,
        { pid: res.pid, state: 'sleep' }
      ]);
      console.log(res.pid,"proceso creadooooooooooooooooo");
    
  } catch (e) {
      console.log(e);
  }
}

const deleteProc = async () => {

   try {
    
    console.log("Proceso a eliminar:", entradaTexto); 
      const req = await deleteProcess(entradaTexto);
      const res = await req.json();
      console.log(res,"Crear Prxxxxxxxxxxxxxxxxxxxxoceso");
      setProcesos((prevProcesos) =>
        prevProcesos.map((proceso) =>
          proceso.pid === parseInt(entradaTexto) ? { ...proceso, state: 'zombie' } : proceso
        )
      );
      setActionButton("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx");
  } catch (e) {
      console.log(e);
  }
  console.log("Proceso a eliminar:", entradaTexto); 
}

  return (
    <>
      
      <div className="miTexto">
      SO1-Proyecto1-2024
    </div>
            <PASTEL />
           
            <div className="card1">
      <button className="my-button"  onClick={createProc}>Crear proceso</button>
      <input 
      type="text" 
      value={entradaTexto} 
      onChange={(e) => setEntradaTexto(e.target.value)} 
      placeholder="Ingrese número de proceso "
    />
    <button className="my-button2"  onClick={deleteProc}>matar proceso</button>
          <p>
              {pid}
          </p>
      </div>
      <TablaProcesos />
      <TablaPid procesos={procesos} />
    </>
  )
}

export default App
