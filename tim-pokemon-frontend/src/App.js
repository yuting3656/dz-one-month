import React from 'react';
import './App.css';
import PokeCard from './components/PokeCard/PokeCard';
import AddPokeCard from './components/AddPokeCard/AddPokeCard';
import {
  Row, 
  Col,
  PageHeader
} from 'antd';
function App() {
  return (
    <div>
      <Row gutter={[8, 8]}>
        <Col span={8}> </Col>
        <Col span={8}>
            <PokeCard/>
        </Col>
        <Col span={8}> </Col>
      </Row>
      <Row>
        <Col>
        <PageHeader title="ADD Pokemon"></PageHeader>
        < AddPokeCard/>
        </Col>
      </Row>
    </div>
  );
}

export default App;
