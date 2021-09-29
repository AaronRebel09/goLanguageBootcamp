package calculadora
// Se importa el package testing
import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)
// se crea un un struct dummyLogger
type dummyLogger struct{}
// Se escriben las funciones necesarios para que dummyLogger cumpla con la interfaz que
// va a reemplazar (Logger)
func (d *dummyLogger) Log(string) error {
	return nil
}
// se crea un struct stubLogger
type stubLogger struct{}
// Se escribe las funciones necesarias para que stubLogger retorne exactamente lo que necesitamos
func (s *stubLogger) Log(string) error {
	return errors.New("error desde stub")
}
// se crea un un struct spy compuesto por un booleano que nos informará si ocurre el llamado a Log
type spyLogger struct {
	spyCalled bool
}
// Para espiar creamos un loggerSpy que setea en true spyCalled si entra al método
func (s *spyLogger) Log(string) error {
	s.spyCalled = true
	return nil
}
// se crea un struct mockConfig
type mockConfig struct {
	clienteUsado string
}
// El mock debe implementar el método necesario y comprobar que SumaEnabled sea llamado y que se
// haga exactamente con el mismo cliente que recibió SumarRestricted
func (m *mockConfig) SumaEnabled(cliente string) bool {
	m.clienteUsado = cliente
	return true
}
// se crea un un struct fakeConfig que implemente una lógica en la que sólo habilita la suma al cliente "John Doe"
type fakeConfig struct{}

func (f *fakeConfig) SumaEnabled(cliente string) bool {
	return cliente == "John Doe"
}

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se ejecuta el test
	resultado := Sumar(num1, num2)
	// mode 1 un complemented
	// Se validan los resultados
	if resultado != resultadoEsperado {
		t.Errorf("Funcion suma() arrojo el resultado = %v" +
			", pero el esperado es %v", resultado, resultadoEsperado)
	}
	// mode 2 complemented with Testify
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
/*	● Validar igualdad -> assert.Equal(t, esperado, obtenido, "mensaje")
	● Desigualdad -> assert.NotEqual(t, esperado, obtenido, "mensaje")
	● Nulo Esperado -> assert.Nil(t, object)
	● No Nulo Esperado -> assert.NotNil(t, object)*/
}

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 5
	num2 := 3
	resultadoEsperado := 2
	// Se ejecuta el test
	resultado := Restar(num1,num2)
	assert.Equal(t, resultadoEsperado, resultado)
}

// Test for demonstarte panic at Test
/*func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 3
	num2 := 0
	// Se ejecuta el test
	resultado := Dividir(num1,num2)
	assert.NotNil(t, resultado)
}*/
func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 6
	num2 := 3
	resultadoEsperado := 2
	// Se ejecuta el test
	resultado := Dividir(num1,num2)
	assert.Equal(t, resultadoEsperado, resultado)
}

func TestSumarLogger(t *testing.T) {
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se genera el objeto dummy a usar para satisfacer la necesidad de la funcion Sumar
	myDummy := &dummyLogger{}
	// se ejecuta el test
	resultado := SumarLogger(num1, num2, myDummy)
	// se validan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado,"deben ser iguales")
}

func TestSumarError(t *testing.T){
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := -99999
	// se genera el objeto stub a usar para satisfacer la necesidad de la funcion Sumar
	myStub := &stubLogger{}
	// se ejecuta el test
	resultado := SumarLogger(num1, num2, myStub)
	// se evaluan los resultados aprovechando testify
	assert.Equal(t, resultadoEsperado, resultado,"deben ser iguales")
}

func TestSumarConSpy(t *testing.T){
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// se genera el objeto spy a usar
	mySpy := &spyLogger{}
	// se ejecuta el test y se validan el resultado y que spyCalledsea true para dar el test valido
	resultado := SumarLogger(num1, num2, mySpy)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.True(t, mySpy.spyCalled)
}

func TestSumarRestricted(t *testing.T) {
	// Se inicializan los datos a usar en el test(input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	resultadoEsperado := 8
	// se genera el objeto dummy a usar para satisfacer la necesidad de la funcion Sumar
	myMock := &mockConfig{}
	// se ejecuta el test y se validan el resultado y que el mock haya registrado la informacion correcta
	resultado := SumarRestricted(num1, num2, myMock, cliente)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	assert.Equal(t, cliente, myMock.clienteUsado)
}

func TestSumarRestrictedFake(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	cliente := "John Doe"
	cliente_dos := "Mister Pmosh"
	resultadoEsperado := 8
	resultadoEsperadoError := -99999
	// Se genera el objeto fake a usar
	myFake := &fakeConfig{}
	// Se ejecuta el test y Se valida que para el cliente autorizado devuelva el resultado
	// correcto de la suma y que para el cliente no autorizado devuelva el número -99999
	resultado := SumarRestricted(num1, num2, myFake, cliente)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
	resultado2 := SumarRestricted(num1, num2, myFake, cliente_dos)
	assert.Equal(t, resultadoEsperadoError, resultado2, "deben ser iguales")
}