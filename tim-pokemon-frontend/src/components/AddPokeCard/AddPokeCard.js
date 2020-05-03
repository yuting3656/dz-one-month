import React from 'react'
import {
    Card
} from 'antd';

import GuitarChord from 'react-guitar-chord'



const AddPokeCard = () => {
    return (
        <Card>
            <GuitarChord chord={'C'} />
        </Card>
    )
}

export default  AddPokeCard;
