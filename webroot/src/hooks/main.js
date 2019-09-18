import { useState, useEffect } from 'react'
import {http} from "../utils/axios"
export function useBusinessList(pageIndex) {
  let [businessList, setBusinessList] = useState([])
  useEffect(() => {
    let url = `/businessList?pageIndex=${pageIndex}&pageSize=10`
    http.get(url).then(({data})=>{
      console.log("data:", data);
      setBusinessList(data.data)
    }, (e)=>{
      console.error(e)
    })
  }, [pageIndex])
  return businessList
}

export function useCount(page) {
  let [count, setCount] = useState(0)
  useEffect(()=>{
    http.get('/count').then(({data})=>{
      setCount(data.data)
    }, e=>console.error(e))
  },[page])
  return count
}