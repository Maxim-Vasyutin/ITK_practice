## Требования  
1. Определите кастомные ошибки:  
 
 ```go
   var (  
       ErrNotFound   = errors.New("ресурс не найден") 
       TimeoutError = errors.New("таймаут операции") 
    )

2. Создайте функцию SimulateRequest() error, которая:  
   - В 50% случаев возвращает TimeoutError, обёрнутую в fmt.Errorf("запрос не выполнен: %w", TimeoutError).  
   - В 30% случаев возвращает ErrNotFound, обёрнутую в fmt.Errorf("ошибка: %w", ErrNotFound).  
   - В 20% случаев возвращает новую ошибку "неизвестная ошибка".  
   - Реализуйте логику анализа ошибок в ProcessError.  
3. Реализуйте логику анализа ошибок в ProcessError.