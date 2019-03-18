describe('Login', function() {
  it('Does not do much!', function() {
    cy.visit('https://localhost');

    cy.url().should('include', '/login');

    cy.get("input[type='email']").type('admin@scores.network');
    cy.get("input[type='password']").type('test123');

    cy.get("button[type='submit']").click();
  });
});
