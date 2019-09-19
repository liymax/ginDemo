import React, {useEffect, useState} from 'react'
import styled from 'styled-components'
import {http} from "../utils/axios"

const Div = styled.div`
  display: flex;
  flex-direction: column;
  font-size: 14px;
  &>em{
    color: blue;
  }
`
const Pre = styled.pre`
  font-size: 15px;
  color: blue;
  font-weight: 600;
`
function SysMemo() {
  let [msg, setMsg] = useState('')
  useEffect(()=>{
    http.get('/sys/memo').then(({data})=>{
      setMsg(data.data)
    })
  })
  /*<Div>{msg.split('\n').map((e,idx)=> <em key={idx+'k'}>{e}</em>)}</Div>*/
  return (
    <Pre>{msg}</Pre>
  )
}

export default SysMemo