import React, { useEffect, useState } from 'react'
import { 
    Card,
    Button
} from 'antd';

import {
    getAllTeamsAPI
} from '../../core/api/teams';

//////////////////////////
////// IMPORT ZONE ///////
//////////////////////////

const PokeCard = () => {
    const [ teams, setTeams ] = useState([])

    useEffect(() => {
        getAllTeamsAPI().then( data => {
            console.log('data', data)
            setTeams(data.data)
        })
    }, [])

    const processTeamsCard = () => 
      teams.map( (data, index) =>  { 
         return ( <Card title={data.name} key={index}>
                 <p>{data.memo}</p>
                 {JSON.stringify(data)}
         </Card> )
          } 
        )

    return (
        <div>
        {processTeamsCard()}
        {teams.map( (data, index) =>  { 
         return <Card title={data.name} key={index}>
                 <p>{data.memo}</p>
         </Card>
          } )
        }
        </div>
    )
}


export default PokeCard;