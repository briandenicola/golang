package m_example

import (
    "testing"
    "m_example"
)

func TestInsert(t *testing.T) {
    t.Log("Inserting People into Collection")
    c := m_example.CreateCollectionSession("localhost:27017", "test", "people")
    
    err := c.Insert( &m_example.Person{"Brian", "(847) 555 1234"} )
    if err != nil {
        t.Error("Received error on import")
    }

    err = c.Insert( &m_example.Person{"George", "(555) 555 5555"} )
    if err != nil {
        t.Error("Received error on import" )
    }
}

func TestQuery(t *testing.T) {
    t.Log("Query People in the Collection")
    c := m_example.CreateCollectionSession("localhost:27017", "test", "people")

    if result := c.Query("Brian"); result.Name != "Brian" {
        t.Errorf("Expected name of 'Brian', but it was %s instead.", result.Name)
    }   
}
