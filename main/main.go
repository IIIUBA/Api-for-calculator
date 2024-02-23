package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Knetic/govaluate"
)

type ArithmeticExpression struct {
	ID               string    `json:"id"`
	ExpressionString string    `json:"expression"`
	State            string    `json:"state"`
	CreationTime     time.Time `json:"creation_time"`
	LastUpdateTime   time.Time `json:"last_update_time"`
	EvaluationResult float64   `json:"evaluation_result"`
	IsEvaluated      bool      `json:"is_evaluated"`
}

type ComputationAgent struct {
	State      string `json:"state"`
	Identifier int    `json:"identifier"`
}

type ExpressionStore struct {
	sync.Mutex
	store map[string]*ArithmeticExpression
}

func NewExpressionStore() *ExpressionStore {
	return &ExpressionStore{
		store: make(map[string]*ArithmeticExpression),
	}
}

func (s *ExpressionStore) Add(expression *ArithmeticExpression) {
	s.Lock()
	defer s.Unlock()
	s.store[expression.ID] = expression
}

func (s *ExpressionStore) GetAll() []*ArithmeticExpression {
	s.Lock()
	defer s.Unlock()
	allExpressions := make([]*ArithmeticExpression, 0, len(s.store))
	for _, expr := range s.store {
		allExpressions = append(allExpressions, expr)
	}
	return allExpressions
}

func (s *ExpressionStore) UpdateState(id string, state string) {
	s.Lock()
	defer s.Unlock()
	if expr, exists := s.store[id]; exists {
		expr.State = state
		expr.LastUpdateTime = time.Now()
	}
}

func (s *ExpressionStore) FetchUnprocessed() *ArithmeticExpression {
	s.Lock()
	defer s.Unlock()
	for _, expr := range s.store {
		if expr.State == "processing" {
			expr.State = "in progress"
			expr.LastUpdateTime = time.Now()
			return expr
		}
	}
	return nil
}

type ComputationManager struct {
	agents          []*ComputationAgent
	processingTime  time.Duration
	expressionStore *ExpressionStore
}

func NewComputationManager(processingTime time.Duration, store *ExpressionStore) *ComputationManager {
	return &ComputationManager{
		processingTime:  processingTime,
		expressionStore: store,
	}
}

func (m *ComputationManager) StartAgents(count int) {
	for i := 0; i < count; i++ {
		agent := &ComputationAgent{
			State:      "idle",
			Identifier: len(m.agents) + 1,
		}
		m.agents = append(m.agents, agent)
		go m.runComputationAgent(agent)
	}
}

func (m *ComputationManager) runComputationAgent(agent *ComputationAgent) {
	for {
		expr := m.expressionStore.FetchUnprocessed()
		if expr == nil {
			agent.State = "idle"
			continue
		}
		agent.State = "working"
		result, err := evaluateExpression(expr.ExpressionString, m.processingTime)
		m.expressionStore.UpdateState(expr.ID, "completed")
		expr.EvaluationResult = result
		expr.IsEvaluated = err == nil
	}
}

func evaluateExpression(exprStr string, processingTime time.Duration) (float64, error) {
	time.Sleep(processingTime)
	expression, err := govaluate.NewEvaluableExpression(exprStr)
	if err != nil {
		return 0, err
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return 0, err
	}
	return result.(float64), nil
}

func main() {
	expressionStore := NewExpressionStore()
	computationManager := NewComputationManager(10*time.Second, expressionStore)

	http.HandleFunc("/expression", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		expression := r.FormValue("expression")
		expression = strings.ReplaceAll(expression, "p", "+")
		expressionID := r.FormValue("id")
		if expressionID == "" {
			expressionID = createUniqueID()
		}
		newExpression := &ArithmeticExpression{
			ID:               expressionID,
			ExpressionString: expression,
			State:            "processing",
			CreationTime:     time.Now(),
			LastUpdateTime:   time.Now(),
			IsEvaluated:      false,
		}
		expressionStore.Add(newExpression)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/expressions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		allExpressions := expressionStore.GetAll()
		json.NewEncoder(w).Encode(allExpressions)
	})

	http.HandleFunc("/computation_agent", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		numberToAdd, err := strconv.Atoi(r.FormValue("add"))
		if err != nil {
			http.Error(w, "Value must be an integer", http.StatusBadRequest)
			return
		}
		computationManager.StartAgents(numberToAdd)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/agents_status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		json.NewEncoder(w).Encode(computationManager.agents)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func createUniqueID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(10000))
}
