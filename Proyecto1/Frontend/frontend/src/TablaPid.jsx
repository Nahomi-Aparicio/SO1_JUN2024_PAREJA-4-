import React, { useState } from 'react';
import './estilos_pagina.css';

const TablaPid = ({ procesos }) => {
  return (
    <div className="tabla-pid-wrapper">
      <table className="tabla-pid">
        <thead>
          <tr>
            <th>PID</th>
            <th>Estado</th>
          </tr>
        </thead>
        <tbody>
          {procesos.map((proceso, index) => (
            <tr key={index}>
              <td>{proceso.pid}</td>
              <td>{proceso.state}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TablaPid;
