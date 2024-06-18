import React, { useState, useEffect } from 'react';
import { Pie } from 'react-chartjs-2';
import { getRam ,getCPU} from './api/Endpoint'; // Asegúrate de tener esta función correctamente importada
import 'chart.js/auto';
import './estilos_pagina.css'; 

const PASTEL = () => {
  const [percentus, setPercentRamus] = useState(0);
  const [percentRamlib, setPercentRamlib] = useState(0);
  const [percentRamtot, setPercentRamtot] = useState(0);


  const [ percentCPU, setPercentcpu ] = useState(0);;


  const [dataRam, setDataRam] = useState({
    labels: ['Usado (GB)', 'Libre (GB)'],
    datasets: [
      {
        data: [0, 0],
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)'],
        hoverOffset: 4,
      },
    ],
  });


  useEffect(() => {
    const interval = setInterval(() => {
      (async () => {
        try {
          const req = await getRam();
          const res = await req.json();
          console.log(res);
          const us = (res.uso / 1073741824).toFixed(2);
          const lib = (res.libre / 1073741824).toFixed(2);
          const tot = (res.total / 1073741824).toFixed(2);
          console.log(us, lib, tot);
          
          setPercentRamus(us);
          setPercentRamlib(lib);
          setPercentRamtot(tot);

          // Actualiza el estado del gráfico
          setDataRam({
            labels: [ 'Libre (%)','Uso (%)'],
            datasets: [
              {
                data: [parseFloat(res.porcent), parseFloat((100-res.porcent))],
                backgroundColor: [ 'rgb(54, 162, 235)','rgb(255, 99, 132)'],
                hoverOffset: 4,
              },
            ],
          });
        } catch (e) {
          console.log(e);
        }
      })();
    }, 500);

    return () => clearInterval(interval);
  }, [percentus],[percentRamlib],[percentRamtot]);


//--------------------------------PARA EL CPU 

  const [dataCpu, setDataCpu] = useState({
    labels: ['Uso (%)', 'Libre (%)'],
    datasets: [
      {
        data: [0, 100],
        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)'],
        hoverOffset: 4,
      },
    ],
  });


  useEffect(() => {
    const interval = setInterval(() => {
        (
            async () => {
                try {
                  const req = await getCPU();
                  const res = await req.json();
                  console.log(res.cpu);
                  setPercentcpu(res.cpu);

                  setDataCpu({
                    labels: ['Uso (%)', 'Libre (%)'],
                    datasets: [
                      {
                        data: [parseFloat(res.cpu), parseFloat((100-res.cpu))],
                        backgroundColor: ['rgb(255, 99, 132)', 'rgb(54, 162, 235)'],
                        hoverOffset: 4,
                      },
                    ],
                  });
                } catch (e) {
                    console.log(e);
                }
            }
        )();
    }, 500);
  
    return () => clearInterval(interval);
  
  }, [percentCPU]);
  






  


  return (
    <>
     

  <div>
  <div className="card" style={{ display: 'flex', justifyContent: 'space-between' }}>
    <Pie data={dataRam} />
    <Pie data={dataCpu} />
  </div>

  <div className="fila">
  <div className="columna">
        <div className="cuadro">
          <div className="titulo">RAM</div>
          <div className="contenido">RAM EN USO :{percentus} GB</div>
          <div className="contenido">RAM LIBRE :{percentRamlib} GB</div>
          <div className="contenido">RAM TOTAL :{percentRamtot} GB</div>
        </div>
      </div>

      <div className="columna">
        
        <div className="cuadro">          
          <div className="titulo">CPU</div>
          <div className="contenido">_</div>
          <div className="contenido">CPU EN USO:{ percentCPU } % </div>
          <div className="contenido">_</div>
          </div>
  </div>
</div>



      </div>
      
  
    </>
  );
};

export default PASTEL;
