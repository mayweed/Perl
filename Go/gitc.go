package main

import (
    "fmt"
    "log"
    )

type graph struct{
    factoryCount int
    linkCount int
    //a slice of int slices(dest + distance)
    edges map[int][][]int
    
    factories []factory
    troops []troop
}
 type factory struct{
     cyborgs int
     production int
     //arg1 to know
     ownership int
 }
 
 type troop struct{
    from int
    to int
    cyborgs int
    remainingTurns int
    ownership int
}
 
func main() {
    // factoryCount: the number of factories
    var factoryCount int
    fmt.Scan(&factoryCount)
    
    // linkCount: the number of links between factories
    var linkCount int
    fmt.Scan(&linkCount)
    
    network:=graph{
        factoryCount:factoryCount,
        linkCount:linkCount,
        edges:make(map[int][][]int),
    }
    
    for i := 0; i < linkCount; i++ {
        var factory1, factory2, distance int
        fmt.Scan(&factory1, &factory2, &distance)
        //not directed, should i use a edge struct?
        network.edges[factory1]=append(network.edges[factory1],[]int{factory2,distance})
        network.edges[factory2]=append(network.edges[factory2],[]int{factory1,distance})
    }
    for {
        // entityCount: the number of entities (e.g. factories and troops)
        var entityCount int
        fmt.Scan(&entityCount)
        
        for i := 0; i < entityCount; i++ {
            var entityId int
            var entityType string
            var arg1, arg2, arg3, arg4, arg5 int
            fmt.Scan(&entityId, &entityType, &arg1, &arg2, &arg3, &arg4, &arg5)
            switch entityType{
                case "FACTORY":
                    network.factories=append(network.factories,factory{arg2,arg3,arg1})
                case "TROOP":
                    network.troops=append(network.troops,troop{arg2,arg3,arg4,arg5,arg1})
            }
        }
        log.Println(network.edges,network.factories,network.troops)
        // Any valid action, such as "WAIT" or "MOVE source destination cyborgs"
        fmt.Println("WAIT")
    }
}