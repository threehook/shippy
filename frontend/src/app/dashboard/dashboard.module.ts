import { NgModule } from '@angular/core';
import { DashboardComponent } from './dashboard.component';
import { MatToolbarModule, MatIconModule, MatButtonModule, MatSidenavModule, MatListModule, MatDividerModule,
  MatMenuModule } from '@angular/material';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [DashboardComponent],
  imports: [MatToolbarModule, MatIconModule, MatButtonModule, MatSidenavModule, RouterModule, MatListModule,
    MatDividerModule, MatMenuModule]
})
export class DashboardModule {}
