import React, { useState, useEffect } from 'react';
import { getCPU } from './api/Endpoint'; // Asegúrate de tener esta función correctamente importada
import './estilos_pagina.css';

const TablaProcesos = () => {
  const [procesos, setProcesos] = useState([]);

  useEffect(() => {
    const interval = setInterval(async () => {
      try {
        const req = await getCPU();
        const res = await req.json();
        console.log(res.Procesos);
        setProcesos(res.procesos);
      } catch (e) {
        console.log(e);
      }
    }, 3500);

    return () => clearInterval(interval);
  }, []);

  const [showHijos, setShowHijos] = useState({});

  const toggleHijos = (index) => {
    setShowHijos((prevState) => ({
      ...prevState,
      [index]: !prevState[index],
    }));
  };

  return (
    <div className="tabla-wrapper">
      <table className="tabla-procesos">
        <thead>
          <tr>
            <th>Proceso</th>
            <th>PID</th>
            <th>Nombre</th>
            <th>Estado</th>
            <th>Mostrar</th>
          </tr>
        </thead>
        <tbody>
          {procesos.map((proceso, index) => (
            <React.Fragment key={index}>
              <tr>
                <td>{`Proceso ${index}`}</td>
                <td>{proceso.pid}</td>
                <td>{proceso.name}</td>
                <td>{proceso.state}</td>
                <td>
                  {proceso.Hijo && proceso.Hijo.length > 0 && (
                    <button onClick={() => toggleHijos(index)}>
                      {showHijos[index] ? 'Ocultar' : 'Mostrar'}
                    </button>
                  )}
                </td>
              </tr>
              {showHijos[index] && proceso.Hijo && proceso.Hijo.length > 0 &&
                proceso.Hijo.map((hijo, hijoIndex) => (
                  <tr key={`${index}-${hijoIndex}`} className="hijo-row">
                    <td>{`Hijo ${hijoIndex}`}</td>
                    <td>{hijo.pid}</td>
                    <td>{hijo.name}</td>
                    <td>{hijo.state}</td>
                    <td></td>
                  </tr>
                ))}
            </React.Fragment>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TablaProcesos;
