import { NgModule } from '@angular/core';
import { MatCardModule, MatButtonModule, MatFormFieldModule, MatInputModule, MatIconModule, MatDividerModule,
  MatToolbarModule, MatDialogModule, MatSelectModule, MatExpansionModule, MatTooltipModule, MatTabsModule,
  MatBadgeModule, MatListModule, MatButtonToggleModule, MatSlideToggleModule, MatDatepickerModule, MatMenuModule,
  MatSnackBarModule } from '@angular/material';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ListDraftComponent } from './list-draft/list-draft.component';
import { HttpClientModule } from '@angular/common/http';
import { ToolboxComponent } from './toolbox/toolbox.component';


@NgModule({
  declarations: [ListDraftComponent, ToolboxComponent],
  providers: [],
  imports: [MatCardModule, MatFormFieldModule, MatDividerModule, HttpClientModule, MatSnackBarModule,
    MatInputModule, MatExpansionModule, MatMenuModule, MatDatepickerModule, MatButtonToggleModule,
    MatSlideToggleModule,  MatBadgeModule, MatTabsModule, MatTooltipModule, MatToolbarModule, MatIconModule,
    MatButtonModule, MatDialogModule, MatSelectModule, FormsModule, ReactiveFormsModule, RouterModule, CommonModule],
  entryComponents: []
})
export class FlexibleAgreementsModule {}
