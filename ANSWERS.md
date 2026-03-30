# ANSWERS

# Question 1

For a FROZEN wallet with `balance=100`, calling `Withdraw(-50)`:

- Version A returns `ErrWalletFrozen` (status check happens first).
- Version B returns `ErrInvalidAmount` (input validation happens first).

Version B is correct for this domain rule set. Input validation should run before state/business checks because `-50` is invalid regardless of wallet state. This gives deterministic validation behavior and clearer API contracts.

# Question 1

When `balance=100` and requested withdrawal is `150`:

- If `Required=150` (requested amount), message can be:  
  `"insufficient balance: required 150, available 100"`
- If `Required=50` (deficit), message can be:  
  `"insufficient balance: short by 50, available 100"`

Better UX: use `Required=150` and `Available=100`. It mirrors the user action ("I tried to withdraw 150") and allows clients to compute deficit when needed (`Required - Available`) without losing requested intent.

# Question 3

Domain errors should not be wrapped in the usecase layer with `fmt.Errorf("failed: %w", err)` because:

- It pollutes domain-level error contracts with transport/usecase text.
- Callers often branch on exact domain errors (`errors.Is` / `errors.As`); wrapping adds unnecessary indirection and can encourage broad, non-domain error handling.
- Domain errors are already meaningful and expected; they should flow unchanged so higher layers can map them cleanly (e.g., to HTTP status, gRPC codes, UI messages).

Wrap only unexpected infrastructure errors where extra operational context is valuable.
